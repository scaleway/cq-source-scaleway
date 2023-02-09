# Table: scaleway_iot_hubs

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|id (PK)|String|
|name|String|
|status|String|
|product_plan|String|
|enabled|Bool|
|device_count|Int|
|connected_device_count|Int|
|endpoint|String|
|disable_events|Bool|
|events_topic_prefix|String|
|region|String|
|created_at|Timestamp|
|updated_at|Timestamp|
|project_id|String|
|organization_id|String|
|enable_device_auto_provisioning|Bool|
|has_custom_ca|Bool|
|twins_graphite_config|JSON|