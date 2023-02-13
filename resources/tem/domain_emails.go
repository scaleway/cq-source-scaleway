package tem

import (
	"context"
	"fmt"
	"time"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/scaleway/cq-source-scaleway/client"
	tem "github.com/scaleway/scaleway-sdk-go/api/tem/v1alpha1"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

func domainEmails() *schema.Table {
	return &schema.Table{
		Name:          "scaleway_tem_emails",
		Resolver:      fetchEmails,
		Transform:     transformers.TransformWithStruct(&tem.Email{}, transformers.WithPrimaryKeys("ID")),
		IsIncremental: true,
	}
}

func fetchEmails(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	p := parent.Item.(*tem.Domain)
	api := tem.NewAPI(cl.SCWClient)

	var startFrom *time.Time

	const key = "tem_emails"
	stateID := cl.ID() + "_" + p.ID

	if cl.Backend != nil {
		value, err := cl.Backend.Get(ctx, key, stateID)
		if err != nil {
			return fmt.Errorf("failed to retrieve state from backend: %w", err)
		}
		if value != "" {
			min, err := time.Parse(time.RFC3339, value)
			if err != nil {
				return fmt.Errorf("retrieved invalid state value: %q %w", value, err)
			}
			startFrom = &min
		}
	}

	limit := uint32(100)
	page := int32(1)

	for {
		response, err := api.ListEmails(&tem.ListEmailsRequest{
			Region:   p.Region,
			DomainID: &p.ID,
			Since:    startFrom,
			PageSize: &limit,
			Page:     &page,
		}, scw.WithContext(ctx))
		if err != nil {
			return err
		}

		for i := range response.Emails {
			if startFrom == nil || (response.Emails[i].CreatedAt != nil && response.Emails[i].CreatedAt.After(*startFrom)) {
				startFrom = response.Emails[i].CreatedAt
			}
		}

		res <- response.Emails
		if len(response.Emails) < int(limit) {
			break
		}

		page++
	}

	if cl.Backend != nil && startFrom != nil {
		return cl.Backend.Set(ctx, key, stateID, startFrom.Format(time.RFC3339))
	}
	return nil
}
