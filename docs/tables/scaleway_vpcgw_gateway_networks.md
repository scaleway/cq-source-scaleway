# Table: scaleway_vpcgw_gateway_networks

This table shows data for Scaleway Vpcgw Gateway Networks.

The primary key for this table is **id**.

## Relations

This table depends on [scaleway_vpcgw_gateways](scaleway_vpcgw_gateways.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|id (PK)|`utf8`|
|created_at|`timestamp[us, tz=UTC]`|
|updated_at|`timestamp[us, tz=UTC]`|
|gateway_id|`utf8`|
|private_network_id|`utf8`|
|mac_address|`utf8`|
|enable_masquerade|`bool`|
|status|`utf8`|
|dhcp|`json`|
|enable_dhcp|`bool`|
|address|`json`|
|zone|`utf8`|