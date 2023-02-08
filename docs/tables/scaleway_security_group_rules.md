# Table: scaleway_security_group_rules

The composite primary key for this table is (**security_group_id**, **id**).

## Relations

This table depends on [scaleway_security_groups](scaleway_security_groups.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|security_group_id (PK)|String|
|id (PK)|String|
|protocol|String|
|direction|String|
|action|String|
|ip_range|JSON|
|dest_port_from|Int|
|dest_port_to|Int|
|position|Int|
|editable|Bool|
|zone|String|