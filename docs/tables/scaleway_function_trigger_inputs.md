# Table: scaleway_function_trigger_inputs

This table shows data for Scaleway Function Trigger Inputs.

The composite primary key for this table is (**trigger_id**, **id**).

## Relations

This table depends on [scaleway_function_triggers](scaleway_function_triggers.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|trigger_id (PK)|`utf8`|
|id (PK)|`utf8`|
|mnq_namespace_id|`utf8`|
|status|`utf8`|
|error_message|`utf8`|
|nats_config|`json`|
|sqs_config|`json`|