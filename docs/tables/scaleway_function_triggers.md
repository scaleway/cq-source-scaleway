# Table: scaleway_function_triggers

This table shows data for Scaleway Function Triggers.

The primary key for this table is **id**.

## Relations

This table depends on [scaleway_functions](scaleway_functions.md).

The following tables depend on scaleway_function_triggers:
  - [scaleway_function_trigger_inputs](scaleway_function_trigger_inputs.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|id (PK)|`utf8`|
|name|`utf8`|
|description|`utf8`|
|type|`utf8`|
|status|`utf8`|
|error_message|`utf8`|
|function_id|`utf8`|
|nats_failure_handling_policy|`json`|
|sqs_failure_handling_policy|`json`|