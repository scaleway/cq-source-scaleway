# Table: scaleway_containers

This table shows data for Scaleway Containers.

The primary key for this table is **id**.

## Relations

The following tables depend on scaleway_containers:
  - [scaleway_container_crons](scaleway_container_crons.md)
  - [scaleway_container_domains](scaleway_container_domains.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|id (PK)|`utf8`|
|name|`utf8`|
|namespace_id|`utf8`|
|status|`utf8`|
|environment_variables|`json`|
|min_scale|`int64`|
|max_scale|`int64`|
|memory_limit|`int64`|
|cpu_limit|`int64`|
|timeout|`json`|
|error_message|`utf8`|
|privacy|`utf8`|
|description|`utf8`|
|registry_image|`utf8`|
|max_concurrency|`int64`|
|domain_name|`utf8`|
|protocol|`utf8`|
|port|`int64`|
|secret_environment_variables|`json`|
|http_option|`utf8`|
|region|`utf8`|