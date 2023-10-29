# Table: scaleway_k8s_nodes

This table shows data for Scaleway K8s Nodes.

The primary key for this table is **id**.

## Relations

This table depends on [scaleway_k8s_clusters](scaleway_k8s_clusters.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|id (PK)|`utf8`|
|pool_id|`utf8`|
|cluster_id|`utf8`|
|provider_id|`utf8`|
|region|`utf8`|
|name|`utf8`|
|public_ip_v4|`inet`|
|public_ip_v6|`inet`|
|conditions|`json`|
|status|`utf8`|
|error_message|`utf8`|
|created_at|`timestamp[us, tz=UTC]`|
|updated_at|`timestamp[us, tz=UTC]`|