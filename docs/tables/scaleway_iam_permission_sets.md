# Table: scaleway_iam_permission_sets

This table shows data for Scaleway IAM Permission Sets.

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|id (PK)|`utf8`|
|name|`utf8`|
|scope_type|`utf8`|
|description|`utf8`|
|categories|`list<item: utf8, nullable>`|