# Table: scaleway_lb_frontend_acls

The composite primary key for this table is (**frontend_id**, **id**).

## Relations

This table depends on [scaleway_lb_frontends](scaleway_lb_frontends.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|frontend_id (PK)|String|
|id (PK)|String|
|name|String|
|match|JSON|
|action|JSON|
|frontend|JSON|
|index|Int|
|created_at|Timestamp|
|updated_at|Timestamp|
|description|String|