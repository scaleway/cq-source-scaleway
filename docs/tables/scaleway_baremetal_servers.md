# Table: scaleway_baremetal_servers

The primary key for this table is **id**.

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
|name|String|
|description|String|
|updated_at|Timestamp|
|created_at|Timestamp|
|status|String|
|offer_id|String|
|offer_name|String|
|tags|StringArray|
|ips|JSON|
|domain|String|
|boot_type|String|
|zone|String|
|install|JSON|
|ping_status|String|
|options|JSON|
|rescue_server|JSON|