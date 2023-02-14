# Table: scaleway_registry_namespaces

The composite primary key for this table is (**id**, **region**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|id (PK)|String|
|name|String|
|description|String|
|organization_id|String|
|project_id|String|
|status|String|
|status_message|String|
|endpoint|String|
|is_public|Bool|
|size|Int|
|created_at|Timestamp|
|updated_at|Timestamp|
|image_count|Int|
|region (PK)|String|