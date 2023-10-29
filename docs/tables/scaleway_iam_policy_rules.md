# Table: scaleway_iam_policy_rules

This table shows data for Scaleway IAM Policy Rules.

The composite primary key for this table is (**policy_id**, **id**).

## Relations

This table depends on [scaleway_iam_policies](scaleway_iam_policies.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|policy_id (PK)|`utf8`|
|id (PK)|`utf8`|
|permission_set_names|`list<item: utf8, nullable>`|
|permission_sets_scope_type|`utf8`|
|project_ids|`list<item: utf8, nullable>`|
|organization_id|`utf8`|
|account_root_user_id|`utf8`|