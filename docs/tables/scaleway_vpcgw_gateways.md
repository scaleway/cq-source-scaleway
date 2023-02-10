# Table: scaleway_vpcgw_gateways

The primary key for this table is **id**.

## Relations

The following tables depend on scaleway_vpcgw_gateways:
  - [scaleway_vpcgw_gateway_networks](scaleway_vpcgw_gateway_networks.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|id (PK)|String|
|organization_id|String|
|project_id|String|
|created_at|Timestamp|
|updated_at|Timestamp|
|type|JSON|
|status|String|
|name|String|
|tags|StringArray|
|ip|JSON|
|gateway_networks|JSON|
|upstream_dns_servers|StringArray|
|version|String|
|can_upgrade_to|String|
|bastion_enabled|Bool|
|bastion_port|Int|
|smtp_enabled|Bool|
|zone|String|