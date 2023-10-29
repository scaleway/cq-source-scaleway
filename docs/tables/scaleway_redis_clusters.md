# Table: scaleway_redis_clusters

This table shows data for Scaleway Redis Clusters.

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|id (PK)|`utf8`|
|name|`utf8`|
|project_id|`utf8`|
|status|`utf8`|
|version|`utf8`|
|endpoints|`json`|
|tags|`list<item: utf8, nullable>`|
|node_type|`utf8`|
|created_at|`timestamp[us, tz=UTC]`|
|updated_at|`timestamp[us, tz=UTC]`|
|tls_enabled|`bool`|
|cluster_settings|`json`|
|acl_rules|`json`|
|cluster_size|`int64`|
|zone|`utf8`|
|user_name|`utf8`|
|upgradable_versions|`list<item: utf8, nullable>`|