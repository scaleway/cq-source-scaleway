# Table: scaleway_function_domains

This table shows data for Scaleway Function Domains.

The primary key for this table is **id**.

## Relations

This table depends on [scaleway_functions](scaleway_functions.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|id (PK)|`utf8`|
|hostname|`utf8`|
|function_id|`utf8`|
|url|`utf8`|
|status|`utf8`|
|error_message|`utf8`|