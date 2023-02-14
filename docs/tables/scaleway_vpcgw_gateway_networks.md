# Table: scaleway_vpcgw_gateway_networks

The primary key for this table is **id**.

## Relations

This table depends on [scaleway_vpcgw_gateways](scaleway_vpcgw_gateways.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|id (PK)|String|
|created_at|Timestamp|
|updated_at|Timestamp|
|gateway_id|String|
|private_network_id|String|
|mac_address|String|
|enable_masquerade|Bool|
|status|String|
|dhcp|JSON|
|enable_dhcp|Bool|
|address|JSON|
|zone|String|