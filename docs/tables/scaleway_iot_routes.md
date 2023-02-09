# Table: scaleway_iot_routes

The composite primary key for this table is (**region**, **id**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|region (PK)|String|
|id (PK)|String|
|name|String|
|hub_id|String|
|topic|String|
|type|String|
|created_at|Timestamp|
|updated_at|Timestamp|