# Table: scaleway_function_triggers

The primary key for this table is **id**.

## Relations

This table depends on [scaleway_functions](scaleway_functions.md).

The following tables depend on scaleway_function_triggers:
  - [scaleway_function_trigger_inputs](scaleway_function_trigger_inputs.md)

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
|type|String|
|status|String|
|error_message|String|
|function_id|String|
|nats_failure_handling_policy|JSON|
|sqs_failure_handling_policy|JSON|