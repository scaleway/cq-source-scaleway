# Table: scaleway_function_crons

This table shows data for Scaleway Function Crons.

The primary key for this table is **id**.

## Relations

This table depends on [scaleway_functions](scaleway_functions.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|id (PK)|`utf8`|
|function_id|`utf8`|
|schedule|`utf8`|
|args|`json`|
|status|`utf8`|
|name|`utf8`|