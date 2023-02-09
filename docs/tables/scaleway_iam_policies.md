# Table: scaleway_iam_policies

The primary key for this table is **id**.

## Relations

The following tables depend on scaleway_iam_policies:
  - [scaleway_iam_policy_rules](scaleway_iam_policy_rules.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|id (PK)|String|
|name|String|
|description|String|
|organization_id|String|
|created_at|Timestamp|
|updated_at|Timestamp|
|editable|Bool|
|nb_rules|Int|
|nb_scopes|Int|
|nb_permission_sets|Int|
|user_id|String|
|group_id|String|
|application_id|String|
|no_principal|Bool|