# Table: scaleway_mnq_namespaces

The primary key for this table is **id**.

## Relations

The following tables depend on scaleway_mnq_namespaces:
  - [scaleway_mnq_namespace_credentials](scaleway_mnq_namespace_credentials.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|id (PK)|String|
|name|String|
|endpoint|String|
|protocol|String|
|project_id|String|
|region|String|
|created_at|Timestamp|
|updated_at|Timestamp|