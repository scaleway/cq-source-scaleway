# Table: scaleway_function_tokens

The primary key for this table is **id**.

## Relations

This table depends on [scaleway_function_namespaces](scaleway_function_namespaces.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|id (PK)|String|
|token|String|
|function_id|String|
|namespace_id|String|
|public_key|String|
|status|String|
|description|String|
|expires_at|Timestamp|