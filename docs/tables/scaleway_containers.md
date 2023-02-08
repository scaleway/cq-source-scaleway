# Table: scaleway_containers

The primary key for this table is **id**.

## Relations

The following tables depend on scaleway_containers:
  - [scaleway_container_crons](scaleway_container_crons.md)
  - [scaleway_container_domains](scaleway_container_domains.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|id (PK)|String|
|name|String|
|namespace_id|String|
|status|String|
|environment_variables|JSON|
|min_scale|Int|
|max_scale|Int|
|memory_limit|Int|
|cpu_limit|Int|
|timeout|JSON|
|error_message|String|
|privacy|String|
|description|String|
|registry_image|String|
|max_concurrency|Int|
|domain_name|String|
|protocol|String|
|port|Int|
|secret_environment_variables|JSON|
|http_option|String|
|region|String|