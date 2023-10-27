# Table: scaleway_rdb_instance_acl_rules

This table shows data for Scaleway Rdb Instance ACL Rules.

The composite primary key for this table is (**ip**, **port**, **protocol**, **direction**, **action**).

## Relations

This table depends on [scaleway_rdb_instances](scaleway_rdb_instances.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|instance_id|`utf8`|
|ip (PK)|`json`|
|port (PK)|`int64`|
|protocol (PK)|`utf8`|
|direction (PK)|`utf8`|
|action (PK)|`utf8`|
|description|`utf8`|