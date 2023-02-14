package account

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/gorilla/mux"
	"github.com/scaleway/cq-source-scaleway/client"
	"github.com/scaleway/scaleway-sdk-go/api/account/v2"
)

func createProjects(router *mux.Router) error {
	var obj account.Project
	if err := faker.FakeObject(&obj); err != nil {
		return err
	}

	router.HandleFunc("/account/v2/projects", func(w http.ResponseWriter, r *http.Request) {
		list := account.ListProjectsResponse{
			Projects:   []*account.Project{&obj},
			TotalCount: 1,
		}

		b, err := json.Marshal(&list)
		if err != nil {
			http.Error(w, "unable to marshal request: "+err.Error(), http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if _, err := w.Write(b); err != nil {
			http.Error(w, "failed to write", http.StatusBadRequest)
			return
		}
	})

	return nil
}

func TestProjects(t *testing.T) {
	client.TestHelper(t, Projects(), createProjects)
}
