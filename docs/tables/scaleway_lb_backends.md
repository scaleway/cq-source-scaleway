# Table: scaleway_lb_backends

This table shows data for Scaleway Lb Backends.

The composite primary key for this table is (**lb_id**, **id**).

## Relations

This table depends on [scaleway_lbs](scaleway_lbs.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|lb_id (PK)|`utf8`|
|id (PK)|`utf8`|
|name|`utf8`|
|forward_protocol|`utf8`|
|forward_port|`int64`|
|forward_port_algorithm|`utf8`|
|sticky_sessions|`utf8`|
|sticky_sessions_cookie_name|`utf8`|
|health_check|`json`|
|pool|`list<item: utf8, nullable>`|
|lb|`json`|
|send_proxy_v2|`bool`|
|timeout_server|`int64`|
|timeout_connect|`int64`|
|timeout_tunnel|`int64`|
|on_marked_down_action|`utf8`|
|proxy_protocol|`utf8`|
|created_at|`timestamp[us, tz=UTC]`|
|updated_at|`timestamp[us, tz=UTC]`|
|failover_host|`utf8`|
|ssl_bridging|`bool`|
|ignore_ssl_server_verify|`bool`|