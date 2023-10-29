# Table: scaleway_baremetal_servers

This table shows data for Scaleway Baremetal Servers.

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|id (PK)|`utf8`|
|organization_id|`utf8`|
|project_id|`utf8`|
|name|`utf8`|
|description|`utf8`|
|updated_at|`timestamp[us, tz=UTC]`|
|created_at|`timestamp[us, tz=UTC]`|
|status|`utf8`|
|offer_id|`utf8`|
|offer_name|`utf8`|
|tags|`list<item: utf8, nullable>`|
|ips|`json`|
|domain|`utf8`|
|boot_type|`utf8`|
|zone|`utf8`|
|install|`json`|
|ping_status|`utf8`|
|options|`json`|
|rescue_server|`json`|