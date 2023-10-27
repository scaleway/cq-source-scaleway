# Table: scaleway_function_tokens

This table shows data for Scaleway Function Tokens.

The primary key for this table is **id**.

## Relations

This table depends on [scaleway_function_namespaces](scaleway_function_namespaces.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|id (PK)|`utf8`|
|token|`utf8`|
|function_id|`utf8`|
|namespace_id|`utf8`|
|public_key|`utf8`|
|status|`utf8`|
|description|`utf8`|
|expires_at|`timestamp[us, tz=UTC]`|