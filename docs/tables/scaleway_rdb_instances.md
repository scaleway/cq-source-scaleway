# Table: scaleway_rdb_instances

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
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|created_at|Timestamp|
|volume|JSON|
|region|String|
|id (PK)|String|
|name|String|
|organization_id|String|
|project_id|String|
|status|String|
|engine|String|
|upgradable_version|JSON|
|endpoint|JSON|
|tags|StringArray|
|settings|JSON|
|backup_schedule|JSON|
|is_ha_cluster|Bool|
|read_replicas|JSON|
|node_type|String|
|init_settings|JSON|
|endpoints|JSON|
|logs_policy|JSON|
|backup_same_region|Bool|
|maintenances|JSON|