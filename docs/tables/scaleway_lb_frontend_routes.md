# Table: scaleway_lb_frontend_routes

The primary key for this table is **id**.

## Relations

This table depends on [scaleway_lb_frontends](scaleway_lb_frontends.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|id (PK)|String|
|frontend_id|String|
|backend_id|String|
|match|JSON|
|created_at|Timestamp|
|updated_at|Timestamp|