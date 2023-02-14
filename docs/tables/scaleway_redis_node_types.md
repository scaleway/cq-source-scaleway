# Table: scaleway_redis_node_types

The composite primary key for this table is (**name**, **zone**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|name (PK)|String|
|stock_status|String|
|description|String|
|vcpus|Int|
|memory|Int|
|disabled|Bool|
|beta|Bool|
|zone (PK)|String|