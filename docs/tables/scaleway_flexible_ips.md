# Table: scaleway_flexible_ips

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
|description|String|
|tags|StringArray|
|updated_at|Timestamp|
|created_at|Timestamp|
|status|String|
|ip_address|JSON|
|mac_address|JSON|
|server_id|String|
|reverse|String|
|zone|String|