# Table: scaleway_lbs

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
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|id (PK)|String|
|name|String|
|description|String|
|status|String|
|instances|JSON|
|organization_id|String|
|project_id|String|
|ip|JSON|
|tags|StringArray|
|frontend_count|Int|
|backend_count|Int|
|type|String|
|subscriber|JSON|
|ssl_compatibility_level|String|
|created_at|Timestamp|
|updated_at|Timestamp|
|private_network_count|Int|
|route_count|Int|
|region|String|
|zone|String|