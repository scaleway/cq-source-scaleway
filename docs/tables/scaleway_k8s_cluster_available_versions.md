# Table: scaleway_k8s_cluster_available_versions

The primary key for this table is **name**.

## Relations

This table depends on [scaleway_k8s_clusters](scaleway_k8s_clusters.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|cluster_id|String|
|name (PK)|String|
|label|String|
|region|String|
|available_cnis|StringArray|
|available_ingresses|StringArray|
|available_container_runtimes|StringArray|
|available_feature_gates|StringArray|
|available_admission_plugins|StringArray|
|available_kubelet_args|JSON|