# Table: scaleway_lb_backend_stats

The composite primary key for this table is (**lb_id**, **instance_id**, **backend_id**, **ip**).

## Relations

This table depends on [scaleway_lbs](scaleway_lbs.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|lb_id (PK)|String|
|instance_id (PK)|String|
|backend_id (PK)|String|
|ip (PK)|String|
|server_state|String|
|server_state_changed_at|Timestamp|
|last_health_check_status|String|