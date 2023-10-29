# Table: scaleway_k8s_clusters

This table shows data for Scaleway K8s Clusters.

The primary key for this table is **id**.

## Relations

The following tables depend on scaleway_k8s_clusters:
  - [scaleway_k8s_cluster_available_versions](scaleway_k8s_cluster_available_versions.md)
  - [scaleway_k8s_nodes](scaleway_k8s_nodes.md)
  - [scaleway_k8s_pools](scaleway_k8s_pools.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|id (PK)|`utf8`|
|type|`utf8`|
|name|`utf8`|
|status|`utf8`|
|version|`utf8`|
|region|`utf8`|
|organization_id|`utf8`|
|project_id|`utf8`|
|tags|`list<item: utf8, nullable>`|
|cni|`utf8`|
|description|`utf8`|
|cluster_url|`utf8`|
|dns_wildcard|`utf8`|
|created_at|`timestamp[us, tz=UTC]`|
|updated_at|`timestamp[us, tz=UTC]`|
|autoscaler_config|`json`|
|dashboard_enabled|`bool`|
|ingress|`utf8`|
|auto_upgrade|`json`|
|upgrade_available|`bool`|
|feature_gates|`list<item: utf8, nullable>`|
|admission_plugins|`list<item: utf8, nullable>`|
|open_id_connect_config|`json`|
|apiserver_cert_sans|`list<item: utf8, nullable>`|