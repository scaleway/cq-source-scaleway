# Table: scaleway_vpcgw_dhcp_entries

This table shows data for Scaleway Vpcgw DHCP Entries.

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|id (PK)|`utf8`|
|created_at|`timestamp[us, tz=UTC]`|
|updated_at|`timestamp[us, tz=UTC]`|
|gateway_network_id|`utf8`|
|mac_address|`utf8`|
|ip_address|`inet`|
|hostname|`utf8`|
|type|`utf8`|
|zone|`utf8`|