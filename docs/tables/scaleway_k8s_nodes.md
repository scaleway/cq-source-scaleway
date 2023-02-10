# Table: scaleway_k8s_nodes

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
|pool_id|String|
|cluster_id|String|
|provider_id|String|
|region|String|
|name|String|
|public_ip_v4|Inet|
|public_ip_v6|Inet|
|conditions|JSON|
|status|String|
|error_message|String|
|created_at|Timestamp|
|updated_at|Timestamp|