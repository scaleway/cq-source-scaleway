# Table: scaleway_iam_policies

This table shows data for Scaleway IAM Policies.

The primary key for this table is **id**.

## Relations

The following tables depend on scaleway_iam_policies:
  - [scaleway_iam_policy_rules](scaleway_iam_policy_rules.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|id (PK)|`utf8`|
|name|`utf8`|
|description|`utf8`|
|organization_id|`utf8`|
|created_at|`timestamp[us, tz=UTC]`|
|updated_at|`timestamp[us, tz=UTC]`|
|editable|`bool`|
|nb_rules|`int64`|
|nb_scopes|`int64`|
|nb_permission_sets|`int64`|
|user_id|`utf8`|
|group_id|`utf8`|
|application_id|`utf8`|
|no_principal|`bool`|