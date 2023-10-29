# Table: scaleway_k8s_versions

This table shows data for Scaleway K8s Versions.

The composite primary key for this table is (**name**, **region**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|name (PK)|`utf8`|
|label|`utf8`|
|region (PK)|`utf8`|
|available_cnis|`list<item: utf8, nullable>`|
|available_ingresses|`list<item: utf8, nullable>`|
|available_container_runtimes|`list<item: utf8, nullable>`|
|available_feature_gates|`list<item: utf8, nullable>`|
|available_admission_plugins|`list<item: utf8, nullable>`|
|available_kubelet_args|`json`|