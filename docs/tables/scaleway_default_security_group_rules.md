# Table: scaleway_default_security_group_rules

This table shows data for Scaleway Default Security Group Rules.

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|id (PK)|`utf8`|
|protocol|`utf8`|
|direction|`utf8`|
|action|`utf8`|
|ip_range|`json`|
|dest_port_from|`int64`|
|dest_port_to|`int64`|
|position|`int64`|
|editable|`bool`|
|zone|`utf8`|