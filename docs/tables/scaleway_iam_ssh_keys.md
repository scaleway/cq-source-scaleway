# Table: scaleway_iam_ssh_keys

This table shows data for Scaleway IAM SSH Keys.

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|id (PK)|`utf8`|
|name|`utf8`|
|public_key|`utf8`|
|fingerprint|`utf8`|
|created_at|`timestamp[us, tz=UTC]`|
|updated_at|`timestamp[us, tz=UTC]`|
|organization_id|`utf8`|
|project_id|`utf8`|
|disabled|`bool`|