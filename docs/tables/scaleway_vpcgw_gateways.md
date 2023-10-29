# Table: scaleway_vpcgw_gateways

This table shows data for Scaleway Vpcgw Gateways.

The primary key for this table is **id**.

## Relations

The following tables depend on scaleway_vpcgw_gateways:
  - [scaleway_vpcgw_gateway_networks](scaleway_vpcgw_gateway_networks.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|id (PK)|`utf8`|
|organization_id|`utf8`|
|project_id|`utf8`|
|created_at|`timestamp[us, tz=UTC]`|
|updated_at|`timestamp[us, tz=UTC]`|
|type|`json`|
|status|`utf8`|
|name|`utf8`|
|tags|`list<item: utf8, nullable>`|
|ip|`json`|
|gateway_networks|`json`|
|upstream_dns_servers|`list<item: utf8, nullable>`|
|version|`utf8`|
|can_upgrade_to|`utf8`|
|bastion_enabled|`bool`|
|bastion_port|`int64`|
|smtp_enabled|`bool`|
|zone|`utf8`|