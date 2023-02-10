# Table: scaleway_registry_image_tags

The primary key for this table is **image_id**.

## Relations

This table depends on [scaleway_registry_images](scaleway_registry_images.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|image_id (PK)|String|
|tags|JSON|