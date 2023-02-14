# Table: scaleway_iam_policy_rules

The composite primary key for this table is (**policy_id**, **id**).

## Relations

This table depends on [scaleway_iam_policies](scaleway_iam_policies.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|policy_id (PK)|String|
|id (PK)|String|
|permission_set_names|StringArray|
|permission_sets_scope_type|String|
|project_ids|StringArray|
|organization_id|String|
|account_root_user_id|String|