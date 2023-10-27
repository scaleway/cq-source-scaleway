# Table: scaleway_iot_hubs

This table shows data for Scaleway IOT Hubs.

The primary key for this table is **id**.

## Relations

The following tables depend on scaleway_iot_hubs:
  - [scaleway_iot_devices](scaleway_iot_devices.md)
  - [scaleway_iot_networks](scaleway_iot_networks.md)
  - [scaleway_iot_routes](scaleway_iot_routes.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|id (PK)|`utf8`|
|name|`utf8`|
|status|`utf8`|
|product_plan|`utf8`|
|enabled|`bool`|
|device_count|`int64`|
|connected_device_count|`int64`|
|endpoint|`utf8`|
|disable_events|`bool`|
|events_topic_prefix|`utf8`|
|region|`utf8`|
|created_at|`timestamp[us, tz=UTC]`|
|updated_at|`timestamp[us, tz=UTC]`|
|project_id|`utf8`|
|organization_id|`utf8`|
|enable_device_auto_provisioning|`bool`|
|has_custom_ca|`bool`|
|twins_graphite_config|`json`|