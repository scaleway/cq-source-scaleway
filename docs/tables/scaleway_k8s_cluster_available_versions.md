# Table: scaleway_k8s_cluster_available_versions

This table shows data for Scaleway K8s Cluster Available Versions.

The primary key for this table is **name**.

## Relations

This table depends on [scaleway_k8s_clusters](scaleway_k8s_clusters.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|cluster_id|`utf8`|
|name (PK)|`utf8`|
|label|`utf8`|
|region|`utf8`|
|available_cnis|`list<item: utf8, nullable>`|
|available_ingresses|`list<item: utf8, nullable>`|
|available_container_runtimes|`list<item: utf8, nullable>`|
|available_feature_gates|`list<item: utf8, nullable>`|
|available_admission_plugins|`list<item: utf8, nullable>`|
|available_kubelet_args|`json`|