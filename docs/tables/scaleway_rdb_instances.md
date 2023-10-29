# Table: scaleway_rdb_instances

This table shows data for Scaleway Rdb Instances.

The primary key for this table is **id**.

## Relations

The following tables depend on scaleway_rdb_instances:
  - [scaleway_rdb_databases](scaleway_rdb_databases.md)
  - [scaleway_rdb_instance_acl_rules](scaleway_rdb_instance_acl_rules.md)
  - [scaleway_rdb_instance_privileges](scaleway_rdb_instance_privileges.md)
  - [scaleway_rdb_users](scaleway_rdb_users.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|created_at|`timestamp[us, tz=UTC]`|
|volume|`json`|
|region|`utf8`|
|id (PK)|`utf8`|
|name|`utf8`|
|organization_id|`utf8`|
|project_id|`utf8`|
|status|`utf8`|
|engine|`utf8`|
|upgradable_version|`json`|
|endpoint|`json`|
|tags|`list<item: utf8, nullable>`|
|settings|`json`|
|backup_schedule|`json`|
|is_ha_cluster|`bool`|
|read_replicas|`json`|
|node_type|`utf8`|
|init_settings|`json`|
|endpoints|`json`|
|logs_policy|`json`|
|backup_same_region|`bool`|
|maintenances|`json`|