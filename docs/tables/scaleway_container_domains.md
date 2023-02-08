# Table: scaleway_container_domains

The primary key for this table is **id**.

## Relations

This table depends on [scaleway_containers](scaleway_containers.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|id (PK)|String|
|hostname|String|
|container_id|String|
|url|String|
|status|String|
|error_message|String|