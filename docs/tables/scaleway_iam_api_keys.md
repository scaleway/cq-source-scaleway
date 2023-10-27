# Table: scaleway_iam_api_keys

This table shows data for Scaleway IAM API Keys.

The primary key for this table is **access_key**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|access_key (PK)|`utf8`|
|secret_key|`utf8`|
|application_id|`utf8`|
|user_id|`utf8`|
|description|`utf8`|
|created_at|`timestamp[us, tz=UTC]`|
|updated_at|`timestamp[us, tz=UTC]`|
|expires_at|`timestamp[us, tz=UTC]`|
|default_project_id|`utf8`|
|editable|`bool`|
|creation_ip|`utf8`|