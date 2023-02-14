# Table: scaleway_k8s_clusters

The primary key for this table is **id**.

## Relations

The following tables depend on scaleway_k8s_clusters:
  - [scaleway_k8s_cluster_available_versions](scaleway_k8s_cluster_available_versions.md)
  - [scaleway_k8s_nodes](scaleway_k8s_nodes.md)
  - [scaleway_k8s_pools](scaleway_k8s_pools.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|id (PK)|String|
|type|String|
|name|String|
|status|String|
|version|String|
|region|String|
|organization_id|String|
|project_id|String|
|tags|StringArray|
|cni|String|
|description|String|
|cluster_url|String|
|dns_wildcard|String|
|created_at|Timestamp|
|updated_at|Timestamp|
|autoscaler_config|JSON|
|dashboard_enabled|Bool|
|ingress|String|
|auto_upgrade|JSON|
|upgrade_available|Bool|
|feature_gates|StringArray|
|admission_plugins|StringArray|
|open_id_connect_config|JSON|
|apiserver_cert_sans|StringArray|