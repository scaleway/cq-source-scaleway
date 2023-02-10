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
	"github.com/scaleway/cq-source-scaleway/resources/iam"
	"github.com/scaleway/cq-source-scaleway/resources/instances"
	"github.com/scaleway/cq-source-scaleway/resources/iot"
	"github.com/scaleway/cq-source-scaleway/resources/ipfs"
	"github.com/scaleway/cq-source-scaleway/resources/k8s"
	"github.com/scaleway/cq-source-scaleway/resources/lb"
	"github.com/scaleway/cq-source-scaleway/resources/marketplace"
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
			applesilicon.ServerTypes(),
			applesilicon.Servers(),
			baremetal.OS(),
			baremetal.Offers(),
			baremetal.Servers(),
			containers.Containers(),
			containers.Namespaces(),
			containers.Tokens(),
			flexibleips.FlexibleIPs(),
			functions.Functions(),
			functions.Namespaces(),
			functions.Runtimes(),
			functions.Tokens(),
			iam.APIKeys(),
			iam.Apps(),
			iam.Groups(),
			iam.PermissionSets(),
			iam.Policies(),
			iam.SSHKeys(),
			iam.Users(),
			instances.DefaultSecurityGroupRules(),
			instances.Images(),
			instances.Instances(),
			instances.SecurityGroups(),
			iot.Devices(),
			iot.Hubs(),
			iot.Networks(),
			iot.Routes(),
			ipfs.Volumes(),
			k8s.Clusters(),
			k8s.Versions(),
			lb.LBs(),
			lb.LBTypes(),
			lb.IPs(),
			lb.Routes(),
			lb.Subscribers(),
			marketplace.Categories(),
			marketplace.Images(),
		},
		client.New,
	)
}
