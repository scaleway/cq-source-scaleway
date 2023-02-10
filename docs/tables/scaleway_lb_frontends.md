# Table: scaleway_lb_frontends

The composite primary key for this table is (**lb_id**, **id**).

## Relations

This table depends on [scaleway_lbs](scaleway_lbs.md).

The following tables depend on scaleway_lb_frontends:
  - [scaleway_lb_frontend_acls](scaleway_lb_frontend_acls.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|lb_id (PK)|String|
|id (PK)|String|
|name|String|
|inbound_port|Int|
|backend|JSON|
|lb|JSON|
|timeout_client|Int|
|certificate|JSON|
|certificate_ids|StringArray|
|created_at|Timestamp|
|updated_at|Timestamp|
|enable_http3|Bool|