# Table: scaleway_vpcgw_dhcps

This table shows data for Scaleway Vpcgw Dhcps.

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|id (PK)|`utf8`|
|organization_id|`utf8`|
|project_id|`utf8`|
|created_at|`timestamp[us, tz=UTC]`|
|updated_at|`timestamp[us, tz=UTC]`|
|subnet|`json`|
|address|`inet`|
|pool_low|`inet`|
|pool_high|`inet`|
|enable_dynamic|`bool`|
|valid_lifetime|`json`|
|renew_timer|`json`|
|rebind_timer|`json`|
|push_default_route|`bool`|
|push_dns_server|`bool`|
|dns_servers_override|`list<item: utf8, nullable>`|
|dns_search|`list<item: utf8, nullable>`|
|dns_local_name|`utf8`|
|zone|`utf8`|