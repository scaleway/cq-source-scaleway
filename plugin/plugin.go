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
	"github.com/scaleway/cq-source-scaleway/resources/mnq"
	"github.com/scaleway/cq-source-scaleway/resources/rdb"
	"github.com/scaleway/cq-source-scaleway/resources/redis"
	"github.com/scaleway/cq-source-scaleway/resources/registry"
	"github.com/scaleway/cq-source-scaleway/resources/secrets"
	"github.com/scaleway/cq-source-scaleway/resources/tem"
	"github.com/scaleway/cq-source-scaleway/resources/test"
	"github.com/scaleway/cq-source-scaleway/resources/vpc"
	"github.com/scaleway/cq-source-scaleway/resources/vpcgw"
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
			flexibleips.FlexibleIPs(),
			functions.Functions(),
			functions.Namespaces(),
			functions.Runtimes(),
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
			iot.Hubs(),
			ipfs.Volumes(),
			k8s.Clusters(),
			k8s.Versions(),
			lb.LBs(),
			lb.LBTypes(),
			lb.IPs(),
			lb.Subscribers(),
			marketplace.Categories(),
			marketplace.Images(),
			mnq.Namespaces(),
			rdb.DatabaseBackups(),
			rdb.DatabaseEngines(),
			rdb.Instances(),
			rdb.Snapshots(),
			rdb.NodeTypes(),
			redis.Clusters(),
			redis.Versions(),
			redis.NodeTypes(),
			registry.Images(),
			registry.Namespaces(),
			secrets.Secrets(),
			tem.Domains(),
			test.Humans(),
			vpc.PrivateNetworks(),
			vpcgw.DHCPs(),
			vpcgw.DHCPEntries(),
			vpcgw.Gateways(),
			vpcgw.IPs(),
		},
		client.New,
	)
}
