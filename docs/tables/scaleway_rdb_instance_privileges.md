# Table: scaleway_rdb_instance_privileges

The composite primary key for this table is (**permission**, **database_name**, **user_name**).

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
|permission (PK)|String|
|database_name (PK)|String|
|user_name (PK)|String|