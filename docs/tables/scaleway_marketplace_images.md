# Table: scaleway_marketplace_images

The primary key for this table is **id**.

## Relations

The following tables depend on scaleway_marketplace_images:
  - [scaleway_marketplace_image_versions](scaleway_marketplace_image_versions.md)

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
|logo|String|
|categories|StringArray|
|created_at|Timestamp|
|updated_at|Timestamp|
|valid_until|Timestamp|
|label|String|