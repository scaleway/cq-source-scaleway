# Table: scaleway_lb_frontend_routes

This table shows data for Scaleway Lb Frontend Routes.

The primary key for this table is **id**.

## Relations

This table depends on [scaleway_lb_frontends](scaleway_lb_frontends.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|id (PK)|`utf8`|
|frontend_id|`utf8`|
|backend_id|`utf8`|
|match|`json`|
|created_at|`timestamp[us, tz=UTC]`|
|updated_at|`timestamp[us, tz=UTC]`|