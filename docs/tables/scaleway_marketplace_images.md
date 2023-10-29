# Table: scaleway_marketplace_images

This table shows data for Scaleway Marketplace Images.

The primary key for this table is **id**.

## Relations

The following tables depend on scaleway_marketplace_images:
  - [scaleway_marketplace_image_versions](scaleway_marketplace_image_versions.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|id (PK)|`utf8`|
|name|`utf8`|
|description|`utf8`|
|logo|`utf8`|
|categories|`list<item: utf8, nullable>`|
|created_at|`timestamp[us, tz=UTC]`|
|updated_at|`timestamp[us, tz=UTC]`|
|valid_until|`timestamp[us, tz=UTC]`|
|label|`utf8`|