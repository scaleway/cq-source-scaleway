# Table: scaleway_security_groups

The primary key for this table is **id**.

## Relations

The following tables depend on scaleway_security_groups:
  - [scaleway_security_group_rules](scaleway_security_group_rules.md)

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
|enable_default_security|Bool|
|inbound_default_policy|String|
|outbound_default_policy|String|
|organization|String|
|project|String|
|tags|StringArray|
|organization_default|Bool|
|project_default|Bool|
|creation_date|Timestamp|
|modification_date|Timestamp|
|servers|JSON|
|stateful|Bool|
|state|String|
|zone|String|