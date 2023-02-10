# Table: scaleway_ipfs_volume_pins

The composite primary key for this table is (**volume_id**, **pin_id**).

## Relations

This table depends on [scaleway_ipfs_volumes](scaleway_ipfs_volumes.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|volume_id (PK)|String|
|pin_id (PK)|String|
|status|String|
|created_at|Timestamp|
|cid|JSON|
|delegates|StringArray|
|info|JSON|