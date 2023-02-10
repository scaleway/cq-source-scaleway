# Table: scaleway_lb_backends

The composite primary key for this table is (**lb_id**, **id**).

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
|id (PK)|String|
|name|String|
|forward_protocol|String|
|forward_port|Int|
|forward_port_algorithm|String|
|sticky_sessions|String|
|sticky_sessions_cookie_name|String|
|health_check|JSON|
|pool|StringArray|
|lb|JSON|
|send_proxy_v2|Bool|
|timeout_server|Int|
|timeout_connect|Int|
|timeout_tunnel|Int|
|on_marked_down_action|String|
|proxy_protocol|String|
|created_at|Timestamp|
|updated_at|Timestamp|
|failover_host|String|
|ssl_bridging|Bool|
|ignore_ssl_server_verify|Bool|