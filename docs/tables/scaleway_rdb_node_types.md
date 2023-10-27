# Table: scaleway_rdb_node_types

This table shows data for Scaleway Rdb Node Types.

The composite primary key for this table is (**name**, **region**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|name (PK)|`utf8`|
|stock_status|`utf8`|
|description|`utf8`|
|vcpus|`int64`|
|memory|`int64`|
|volume_constraint|`json`|
|is_bssd_compatible|`bool`|
|disabled|`bool`|
|beta|`bool`|
|available_volume_types|`json`|
|is_ha_required|`bool`|
|region (PK)|`utf8`|