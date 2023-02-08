# Table: scaleway_functions

The primary key for this table is **id**.

## Relations

The following tables depend on scaleway_functions:
  - [scaleway_function_crons](scaleway_function_crons.md)
  - [scaleway_function_domains](scaleway_function_domains.md)
  - [scaleway_function_triggers](scaleway_function_triggers.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|id (PK)|String|
|name|String|
|namespace_id|String|
|status|String|
|environment_variables|JSON|
|min_scale|Int|
|max_scale|Int|
|runtime|String|
|memory_limit|Int|
|cpu_limit|Int|
|timeout|JSON|
|handler|String|
|error_message|String|
|privacy|String|
|description|String|
|domain_name|String|
|secret_environment_variables|JSON|
|region|String|
|http_option|String|
|runtime_message|String|