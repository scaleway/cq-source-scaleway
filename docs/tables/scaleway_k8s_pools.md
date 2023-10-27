# Table: scaleway_k8s_pools

This table shows data for Scaleway K8s Pools.

The primary key for this table is **id**.

## Relations

This table depends on [scaleway_k8s_clusters](scaleway_k8s_clusters.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|id (PK)|`utf8`|
|cluster_id|`utf8`|
|created_at|`timestamp[us, tz=UTC]`|
|updated_at|`timestamp[us, tz=UTC]`|
|name|`utf8`|
|status|`utf8`|
|version|`utf8`|
|node_type|`utf8`|
|autoscaling|`bool`|
|size|`int64`|
|min_size|`int64`|
|max_size|`int64`|
|container_runtime|`utf8`|
|autohealing|`bool`|
|tags|`list<item: utf8, nullable>`|
|placement_group_id|`utf8`|
|kubelet_args|`json`|
|upgrade_policy|`json`|
|zone|`utf8`|
|root_volume_type|`utf8`|
|root_volume_size|`int64`|
|region|`utf8`|