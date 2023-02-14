# Table: scaleway_ipfs_volumes

The primary key for this table is **id**.

## Relations

The following tables depend on scaleway_ipfs_volumes:
  - [scaleway_ipfs_volume_pins](scaleway_ipfs_volume_pins.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|id (PK)|String|
|project_id|String|
|region|String|
|count_pin|Int|
|created_at|Timestamp|
|updated_at|Timestamp|
|tags|StringArray|
|name|String|