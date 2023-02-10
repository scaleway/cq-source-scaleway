# Table: scaleway_k8s_pools

The primary key for this table is **id**.

## Relations

This table depends on [scaleway_k8s_clusters](scaleway_k8s_clusters.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|id (PK)|String|
|cluster_id|String|
|created_at|Timestamp|
|updated_at|Timestamp|
|name|String|
|status|String|
|version|String|
|node_type|String|
|autoscaling|Bool|
|size|Int|
|min_size|Int|
|max_size|Int|
|container_runtime|String|
|autohealing|Bool|
|tags|StringArray|
|placement_group_id|String|
|kubelet_args|JSON|
|upgrade_policy|JSON|
|zone|String|
|root_volume_type|String|
|root_volume_size|Int|
|region|String|