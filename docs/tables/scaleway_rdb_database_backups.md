# Table: scaleway_rdb_database_backups

This table shows data for Scaleway Rdb Database Backups.

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|id (PK)|`utf8`|
|instance_id|`utf8`|
|database_name|`utf8`|
|name|`utf8`|
|status|`utf8`|
|size|`int64`|
|expires_at|`timestamp[us, tz=UTC]`|
|created_at|`timestamp[us, tz=UTC]`|
|updated_at|`timestamp[us, tz=UTC]`|
|instance_name|`utf8`|
|download_url|`utf8`|
|download_url_expires_at|`timestamp[us, tz=UTC]`|
|region|`utf8`|
|same_region|`bool`|