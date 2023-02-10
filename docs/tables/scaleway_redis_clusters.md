# Table: scaleway_redis_clusters

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|id (PK)|String|
|name|String|
|project_id|String|
|status|String|
|version|String|
|endpoints|JSON|
|tags|StringArray|
|node_type|String|
|created_at|Timestamp|
|updated_at|Timestamp|
|tls_enabled|Bool|
|cluster_settings|JSON|
|acl_rules|JSON|
|cluster_size|Int|
|zone|String|
|user_name|String|
|upgradable_versions|StringArray|