# Table: scaleway_rdb_databases

The primary key for this table is **name**.

## Relations

This table depends on [scaleway_rdb_instances](scaleway_rdb_instances.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|instance_id|String|
|name (PK)|String|
|owner|String|
|managed|Bool|
|size|Int|