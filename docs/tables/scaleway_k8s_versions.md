# Table: scaleway_k8s_versions

The composite primary key for this table is (**name**, **region**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|name (PK)|String|
|label|String|
|region (PK)|String|
|available_cnis|StringArray|
|available_ingresses|StringArray|
|available_container_runtimes|StringArray|
|available_feature_gates|StringArray|
|available_admission_plugins|StringArray|
|available_kubelet_args|JSON|