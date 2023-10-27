# Table: scaleway_ipfs_volumes

This table shows data for Scaleway Ipfs Volumes.

The primary key for this table is **id**.

## Relations

The following tables depend on scaleway_ipfs_volumes:
  - [scaleway_ipfs_volume_pins](scaleway_ipfs_volume_pins.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|id (PK)|`utf8`|
|project_id|`utf8`|
|region|`utf8`|
|count_pin|`int64`|
|created_at|`timestamp[us, tz=UTC]`|
|updated_at|`timestamp[us, tz=UTC]`|
|tags|`list<item: utf8, nullable>`|
|name|`utf8`|