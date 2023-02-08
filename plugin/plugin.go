package plugin

import (
	"github.com/cloudquery/plugin-sdk/plugins/source"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/scaleway/cq-source-scaleway/client"
	"github.com/scaleway/cq-source-scaleway/resources/account"
	"github.com/scaleway/cq-source-scaleway/resources/applesilicon"
	"github.com/scaleway/cq-source-scaleway/resources/baremetal"
	"github.com/scaleway/cq-source-scaleway/resources/containers"
	"github.com/scaleway/cq-source-scaleway/resources/flexibleips"
	"github.com/scaleway/cq-source-scaleway/resources/functions"
	"github.com/scaleway/cq-source-scaleway/resources/instances"
)

var (
	Version = "development"
)

func Plugin() *source.Plugin {
	return source.NewPlugin(
		"scaleway",
		Version,
		schema.Tables{
			account.Projects(),
			applesilicon.OS(),
			applesilicon.Servers(),
			applesilicon.ServerTypes(),
			baremetal.Offers(),
			baremetal.OS(),
			baremetal.Servers(),
			containers.Containers(),
			containers.Namespaces(),
			containers.Tokens(),
			flexibleips.FlexibleIPs(),
			functions.Functions(),
			functions.Namespaces(),
			functions.Runtimes(),
			functions.Tokens(),
			instances.DefaultSecurityGroupRules(),
			instances.Images(),
			instances.Instances(),
			instances.SecurityGroups(),
		},
		client.New,
	)
}
