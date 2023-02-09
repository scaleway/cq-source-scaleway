# Table: scaleway_iam_api_keys

The primary key for this table is **access_key**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|access_key (PK)|String|
|secret_key|String|
|application_id|String|
|user_id|String|
|description|String|
|created_at|Timestamp|
|updated_at|Timestamp|
|expires_at|Timestamp|
|default_project_id|String|
|editable|Bool|
|creation_ip|String|