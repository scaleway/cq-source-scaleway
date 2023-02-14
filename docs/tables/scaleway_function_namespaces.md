# Table: scaleway_function_namespaces

The primary key for this table is **id**.

## Relations

The following tables depend on scaleway_function_namespaces:
  - [scaleway_function_tokens](scaleway_function_tokens.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|id (PK)|String|
|name|String|
|environment_variables|JSON|
|organization_id|String|
|project_id|String|
|status|String|
|registry_namespace_id|String|
|error_message|String|
|registry_endpoint|String|
|description|String|
|secret_environment_variables|JSON|
|region|String|