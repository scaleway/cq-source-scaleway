# Table: scaleway_functions

This table shows data for Scaleway Functions.

The primary key for this table is **id**.

## Relations

The following tables depend on scaleway_functions:
  - [scaleway_function_crons](scaleway_function_crons.md)
  - [scaleway_function_domains](scaleway_function_domains.md)
  - [scaleway_function_triggers](scaleway_function_triggers.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|id (PK)|`utf8`|
|name|`utf8`|
|namespace_id|`utf8`|
|status|`utf8`|
|environment_variables|`json`|
|min_scale|`int64`|
|max_scale|`int64`|
|runtime|`utf8`|
|memory_limit|`int64`|
|cpu_limit|`int64`|
|timeout|`json`|
|handler|`utf8`|
|error_message|`utf8`|
|privacy|`utf8`|
|description|`utf8`|
|domain_name|`utf8`|
|secret_environment_variables|`json`|
|region|`utf8`|
|http_option|`utf8`|
|runtime_message|`utf8`|