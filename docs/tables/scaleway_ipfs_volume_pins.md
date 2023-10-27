# Table: scaleway_ipfs_volume_pins

This table shows data for Scaleway Ipfs Volume Pins.

The composite primary key for this table is (**volume_id**, **pin_id**).

## Relations

This table depends on [scaleway_ipfs_volumes](scaleway_ipfs_volumes.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|volume_id (PK)|`utf8`|
|pin_id (PK)|`utf8`|
|status|`utf8`|
|created_at|`timestamp[us, tz=UTC]`|
|cid|`json`|
|delegates|`list<item: utf8, nullable>`|
|info|`json`|