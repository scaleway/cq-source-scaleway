# Table: scaleway_container_tokens

The primary key for this table is **id**.

## Relations

This table depends on [scaleway_container_namespaces](scaleway_container_namespaces.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|id (PK)|String|
|token|String|
|container_id|String|
|namespace_id|String|
|public_key|String|
|status|String|
|description|String|
|expires_at|Timestamp|