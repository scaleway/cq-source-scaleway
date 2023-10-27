# Table: scaleway_registry_images

This table shows data for Scaleway Registry Images.

The primary key for this table is **id**.

## Relations

The following tables depend on scaleway_registry_images:
  - [scaleway_registry_image_tags](scaleway_registry_image_tags.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|id (PK)|`utf8`|
|name|`utf8`|
|namespace_id|`utf8`|
|status|`utf8`|
|status_message|`utf8`|
|visibility|`utf8`|
|size|`int64`|
|created_at|`timestamp[us, tz=UTC]`|
|updated_at|`timestamp[us, tz=UTC]`|
|tags|`list<item: utf8, nullable>`|