# Table: scaleway_registry_image_tags

This table shows data for Scaleway Registry Image Tags.

The primary key for this table is **image_id**.

## Relations

This table depends on [scaleway_registry_images](scaleway_registry_images.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|image_id (PK)|`utf8`|
|tags|`json`|