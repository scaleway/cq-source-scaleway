# Table: scaleway_marketplace_image_versions

This table shows data for Scaleway Marketplace Image Versions.

The primary key for this table is **id**.

## Relations

This table depends on [scaleway_marketplace_images](scaleway_marketplace_images.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|image_id|`utf8`|
|id (PK)|`utf8`|
|name|`utf8`|
|created_at|`timestamp[us, tz=UTC]`|
|updated_at|`timestamp[us, tz=UTC]`|
|published_at|`timestamp[us, tz=UTC]`|