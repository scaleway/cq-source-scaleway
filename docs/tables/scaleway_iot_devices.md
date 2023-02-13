# Table: scaleway_iot_devices

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
|description|String|
|status|String|
|hub_id|String|
|last_activity_at|Timestamp|
|is_connected|Bool|
|allow_insecure|Bool|
|allow_multiple_connections|Bool|
|message_filters|JSON|
|has_custom_certificate|Bool|
|created_at|Timestamp|
|updated_at|Timestamp|