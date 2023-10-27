# Table: scaleway_rdb_instance_privileges

This table shows data for Scaleway Rdb Instance Privileges.

The composite primary key for this table is (**permission**, **database_name**, **user_name**).

## Relations

This table depends on [scaleway_rdb_instances](scaleway_rdb_instances.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|instance_id|`utf8`|
|permission (PK)|`utf8`|
|database_name (PK)|`utf8`|
|user_name (PK)|`utf8`|