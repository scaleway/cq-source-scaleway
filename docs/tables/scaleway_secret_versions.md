# Table: scaleway_secret_versions

The composite primary key for this table is (**secret_id**, **revision**).

## Relations

This table depends on [scaleway_secrets](scaleway_secrets.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|secret_id (PK)|String|
|revision (PK)|Int|
|status|String|
|created_at|Timestamp|
|updated_at|Timestamp|
|description|String|