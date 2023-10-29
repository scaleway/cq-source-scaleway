# Table: scaleway_security_group_rules

This table shows data for Scaleway Security Group Rules.

The composite primary key for this table is (**security_group_id**, **id**).

## Relations

This table depends on [scaleway_security_groups](scaleway_security_groups.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|security_group_id (PK)|`utf8`|
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