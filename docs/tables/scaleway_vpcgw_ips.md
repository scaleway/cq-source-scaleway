# Table: scaleway_vpcgw_ips

This table shows data for Scaleway Vpcgw IPs.

The primary key for this table is **id**.

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
|tags|`list<item: utf8, nullable>`|
|address|`inet`|
|reverse|`utf8`|
|gateway_id|`utf8`|
|zone|`utf8`|