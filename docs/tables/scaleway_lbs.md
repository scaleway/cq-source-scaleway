# Table: scaleway_lbs

This table shows data for Scaleway Lbs.

The primary key for this table is **id**.

## Relations

The following tables depend on scaleway_lbs:
  - [scaleway_lb_backend_stats](scaleway_lb_backend_stats.md)
  - [scaleway_lb_backends](scaleway_lb_backends.md)
  - [scaleway_lb_certificates](scaleway_lb_certificates.md)
  - [scaleway_lb_frontends](scaleway_lb_frontends.md)
  - [scaleway_lb_private_networks](scaleway_lb_private_networks.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|id (PK)|`utf8`|
|name|`utf8`|
|description|`utf8`|
|status|`utf8`|
|instances|`json`|
|organization_id|`utf8`|
|project_id|`utf8`|
|ip|`json`|
|tags|`list<item: utf8, nullable>`|
|frontend_count|`int64`|
|backend_count|`int64`|
|type|`utf8`|
|subscriber|`json`|
|ssl_compatibility_level|`utf8`|
|created_at|`timestamp[us, tz=UTC]`|
|updated_at|`timestamp[us, tz=UTC]`|
|private_network_count|`int64`|
|route_count|`int64`|
|region|`utf8`|
|zone|`utf8`|