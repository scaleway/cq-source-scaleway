# Table: scaleway_applesilicon_server_types

The composite primary key for this table is (**zone**, **name**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|zone (PK)|String|
|cpu|JSON|
|disk|JSON|
|name (PK)|String|
|memory|JSON|
|stock|String|
|minimum_lease_duration|JSON|