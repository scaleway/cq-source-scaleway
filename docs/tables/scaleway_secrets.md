# Table: scaleway_secrets

This table shows data for Scaleway Secrets.

The primary key for this table is **id**.

## Relations

The following tables depend on scaleway_secrets:
  - [scaleway_secret_versions](scaleway_secret_versions.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|id (PK)|`utf8`|
|project_id|`utf8`|
|name|`utf8`|
|status|`utf8`|
|created_at|`timestamp[us, tz=UTC]`|
|updated_at|`timestamp[us, tz=UTC]`|
|tags|`list<item: utf8, nullable>`|
|region|`utf8`|
|version_count|`int64`|
|description|`utf8`|