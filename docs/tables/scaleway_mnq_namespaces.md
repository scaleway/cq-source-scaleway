# Table: scaleway_mnq_namespaces

This table shows data for Scaleway Mnq Namespaces.

The primary key for this table is **id**.

## Relations

The following tables depend on scaleway_mnq_namespaces:
  - [scaleway_mnq_namespace_credentials](scaleway_mnq_namespace_credentials.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|id (PK)|`utf8`|
|name|`utf8`|
|endpoint|`utf8`|
|protocol|`utf8`|
|project_id|`utf8`|
|region|`utf8`|
|created_at|`timestamp[us, tz=UTC]`|
|updated_at|`timestamp[us, tz=UTC]`|