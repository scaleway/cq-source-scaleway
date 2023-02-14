# Table: scaleway_registry_images

The primary key for this table is **id**.

## Relations

The following tables depend on scaleway_registry_images:
  - [scaleway_registry_image_tags](scaleway_registry_image_tags.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|id (PK)|String|
|name|String|
|namespace_id|String|
|status|String|
|status_message|String|
|visibility|String|
|size|Int|
|created_at|Timestamp|
|updated_at|Timestamp|
|tags|StringArray|