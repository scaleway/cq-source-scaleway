# Table: scaleway_security_groups

This table shows data for Scaleway Security Groups.

The primary key for this table is **id**.

## Relations

The following tables depend on scaleway_security_groups:
  - [scaleway_security_group_rules](scaleway_security_group_rules.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|id (PK)|`utf8`|
|name|`utf8`|
|description|`utf8`|
|enable_default_security|`bool`|
|inbound_default_policy|`utf8`|
|outbound_default_policy|`utf8`|
|organization|`utf8`|
|project|`utf8`|
|tags|`list<item: utf8, nullable>`|
|organization_default|`bool`|
|project_default|`bool`|
|creation_date|`timestamp[us, tz=UTC]`|
|modification_date|`timestamp[us, tz=UTC]`|
|servers|`json`|
|stateful|`bool`|
|state|`utf8`|
|zone|`utf8`|