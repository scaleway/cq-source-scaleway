# Table: scaleway_redis_node_types

This table shows data for Scaleway Redis Node Types.

The composite primary key for this table is (**name**, **zone**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|name (PK)|`utf8`|
|stock_status|`utf8`|
|description|`utf8`|
|vcpus|`int64`|
|memory|`int64`|
|disabled|`bool`|
|beta|`bool`|
|zone (PK)|`utf8`|