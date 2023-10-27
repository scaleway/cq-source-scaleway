# Table: scaleway_container_namespaces

This table shows data for Scaleway Container Namespaces.

The primary key for this table is **id**.

## Relations

The following tables depend on scaleway_container_namespaces:
  - [scaleway_container_tokens](scaleway_container_tokens.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|id (PK)|`utf8`|
|name|`utf8`|
|environment_variables|`json`|
|organization_id|`utf8`|
|project_id|`utf8`|
|status|`utf8`|
|registry_namespace_id|`utf8`|
|error_message|`utf8`|
|registry_endpoint|`utf8`|
|description|`utf8`|
|secret_environment_variables|`json`|
|region|`utf8`|