# Table: scaleway_redis_versions

The composite primary key for this table is (**version**, **zone**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|version (PK)|String|
|end_of_life_at|Timestamp|
|available_settings|JSON|
|logo_url|String|
|zone (PK)|String|