# Table: scaleway_iot_routes

The primary key for this table is **id**.

## Relations

This table depends on [scaleway_iot_hubs](scaleway_iot_hubs.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|id (PK)|String|
|name|String|
|hub_id|String|
|topic|String|
|type|String|
|created_at|Timestamp|
|updated_at|Timestamp|