# Table: scaleway_function_trigger_inputs

The composite primary key for this table is (**trigger_id**, **id**).

## Relations

This table depends on [scaleway_function_triggers](scaleway_function_triggers.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|trigger_id (PK)|String|
|id (PK)|String|
|mnq_namespace_id|String|
|status|String|
|error_message|String|
|nats_config|JSON|
|sqs_config|JSON|