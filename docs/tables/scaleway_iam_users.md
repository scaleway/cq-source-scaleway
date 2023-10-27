# Table: scaleway_iam_users

This table shows data for Scaleway IAM Users.

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|id (PK)|`utf8`|
|email|`utf8`|
|created_at|`timestamp[us, tz=UTC]`|
|updated_at|`timestamp[us, tz=UTC]`|
|organization_id|`utf8`|
|deletable|`bool`|
|last_login_at|`timestamp[us, tz=UTC]`|
|type|`utf8`|
|two_factor_enabled|`bool`|
|status|`utf8`|