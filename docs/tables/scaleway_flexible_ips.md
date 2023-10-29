# Table: scaleway_flexible_ips

This table shows data for Scaleway Flexible IPs.

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|id (PK)|`utf8`|
|organization_id|`utf8`|
|project_id|`utf8`|
|description|`utf8`|
|tags|`list<item: utf8, nullable>`|
|updated_at|`timestamp[us, tz=UTC]`|
|created_at|`timestamp[us, tz=UTC]`|
|status|`utf8`|
|ip_address|`json`|
|mac_address|`json`|
|server_id|`utf8`|
|reverse|`utf8`|
|zone|`utf8`|