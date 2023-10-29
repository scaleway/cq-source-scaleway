# Table: scaleway_redis_versions

This table shows data for Scaleway Redis Versions.

The composite primary key for this table is (**version**, **zone**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|version (PK)|`utf8`|
|end_of_life_at|`timestamp[us, tz=UTC]`|
|available_settings|`json`|
|logo_url|`utf8`|
|zone (PK)|`utf8`|