# Table: scaleway_secrets

The primary key for this table is **id**.

## Relations

The following tables depend on scaleway_secrets:
  - [scaleway_secret_versions](scaleway_secret_versions.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|id (PK)|String|
|project_id|String|
|name|String|
|status|String|
|created_at|Timestamp|
|updated_at|Timestamp|
|tags|StringArray|
|region|String|
|version_count|Int|
|description|String|