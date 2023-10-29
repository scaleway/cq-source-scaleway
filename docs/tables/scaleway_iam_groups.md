# Table: scaleway_iam_groups

This table shows data for Scaleway IAM Groups.

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|id (PK)|`utf8`|
|created_at|`timestamp[us, tz=UTC]`|
|updated_at|`timestamp[us, tz=UTC]`|
|organization_id|`utf8`|
|name|`utf8`|
|description|`utf8`|
|user_ids|`list<item: utf8, nullable>`|
|application_ids|`list<item: utf8, nullable>`|