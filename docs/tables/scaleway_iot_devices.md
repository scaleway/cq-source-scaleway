# Table: scaleway_iot_devices

This table shows data for Scaleway IOT Devices.

The primary key for this table is **id**.

## Relations

This table depends on [scaleway_iot_hubs](scaleway_iot_hubs.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|id (PK)|`utf8`|
|name|`utf8`|
|description|`utf8`|
|status|`utf8`|
|hub_id|`utf8`|
|last_activity_at|`timestamp[us, tz=UTC]`|
|is_connected|`bool`|
|allow_insecure|`bool`|
|allow_multiple_connections|`bool`|
|message_filters|`json`|
|has_custom_certificate|`bool`|
|created_at|`timestamp[us, tz=UTC]`|
|updated_at|`timestamp[us, tz=UTC]`|