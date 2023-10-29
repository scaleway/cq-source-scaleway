# Table: scaleway_container_tokens

This table shows data for Scaleway Container Tokens.

The primary key for this table is **id**.

## Relations

This table depends on [scaleway_container_namespaces](scaleway_container_namespaces.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|id (PK)|`utf8`|
|token|`utf8`|
|container_id|`utf8`|
|namespace_id|`utf8`|
|public_key|`utf8`|
|status|`utf8`|
|description|`utf8`|
|expires_at|`timestamp[us, tz=UTC]`|