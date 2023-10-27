# Table: scaleway_lb_frontends

This table shows data for Scaleway Lb Frontends.

The composite primary key for this table is (**lb_id**, **id**).

## Relations

This table depends on [scaleway_lbs](scaleway_lbs.md).

The following tables depend on scaleway_lb_frontends:
  - [scaleway_lb_frontend_acls](scaleway_lb_frontend_acls.md)
  - [scaleway_lb_frontend_routes](scaleway_lb_frontend_routes.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|lb_id (PK)|`utf8`|
|id (PK)|`utf8`|
|name|`utf8`|
|inbound_port|`int64`|
|backend|`json`|
|lb|`json`|
|timeout_client|`int64`|
|certificate|`json`|
|certificate_ids|`list<item: utf8, nullable>`|
|created_at|`timestamp[us, tz=UTC]`|
|updated_at|`timestamp[us, tz=UTC]`|
|enable_http3|`bool`|