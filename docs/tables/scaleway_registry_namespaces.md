# Table: scaleway_registry_namespaces

This table shows data for Scaleway Registry Namespaces.

The composite primary key for this table is (**id**, **region**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|id (PK)|`utf8`|
|name|`utf8`|
|description|`utf8`|
|organization_id|`utf8`|
|project_id|`utf8`|
|status|`utf8`|
|status_message|`utf8`|
|endpoint|`utf8`|
|is_public|`bool`|
|size|`int64`|
|created_at|`timestamp[us, tz=UTC]`|
|updated_at|`timestamp[us, tz=UTC]`|
|image_count|`int64`|
|region (PK)|`utf8`|