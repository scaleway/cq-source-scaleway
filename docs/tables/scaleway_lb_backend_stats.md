# Table: scaleway_lb_backend_stats

This table shows data for Scaleway Lb Backend Stats.

The composite primary key for this table is (**lb_id**, **instance_id**, **backend_id**, **ip**).

## Relations

This table depends on [scaleway_lbs](scaleway_lbs.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|lb_id (PK)|`utf8`|
|instance_id (PK)|`utf8`|
|backend_id (PK)|`utf8`|
|ip (PK)|`utf8`|
|server_state|`utf8`|
|server_state_changed_at|`timestamp[us, tz=UTC]`|
|last_health_check_status|`utf8`|