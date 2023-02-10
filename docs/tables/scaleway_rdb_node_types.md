# Table: scaleway_rdb_node_types

The composite primary key for this table is (**name**, **region**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|name (PK)|String|
|stock_status|String|
|description|String|
|vcpus|Int|
|memory|Int|
|volume_constraint|JSON|
|is_bssd_compatible|Bool|
|disabled|Bool|
|beta|Bool|
|available_volume_types|JSON|
|is_ha_required|Bool|
|region (PK)|String|