# Table: scaleway_default_security_group_rules

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
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