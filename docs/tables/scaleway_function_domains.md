# Table: scaleway_function_domains

The primary key for this table is **id**.

## Relations

This table depends on [scaleway_functions](scaleway_functions.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|id (PK)|String|
|hostname|String|
|function_id|String|
|url|String|
|status|String|
|error_message|String|