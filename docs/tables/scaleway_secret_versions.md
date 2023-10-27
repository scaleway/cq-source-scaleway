# Table: scaleway_secret_versions

This table shows data for Scaleway Secret Versions.

The composite primary key for this table is (**secret_id**, **revision**).

## Relations

This table depends on [scaleway_secrets](scaleway_secrets.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|secret_id (PK)|`utf8`|
|revision (PK)|`int64`|
|status|`utf8`|
|created_at|`timestamp[us, tz=UTC]`|
|updated_at|`timestamp[us, tz=UTC]`|
|description|`utf8`|