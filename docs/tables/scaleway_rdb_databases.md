# Table: scaleway_rdb_databases

This table shows data for Scaleway Rdb Databases.

The primary key for this table is **name**.

## Relations

This table depends on [scaleway_rdb_instances](scaleway_rdb_instances.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|instance_id|`utf8`|
|name (PK)|`utf8`|
|owner|`utf8`|
|managed|`bool`|
|size|`int64`|