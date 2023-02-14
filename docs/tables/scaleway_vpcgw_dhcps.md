# Table: scaleway_vpcgw_dhcps

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|id (PK)|String|
|organization_id|String|
|project_id|String|
|created_at|Timestamp|
|updated_at|Timestamp|
|subnet|JSON|
|address|Inet|
|pool_low|Inet|
|pool_high|Inet|
|enable_dynamic|Bool|
|valid_lifetime|JSON|
|renew_timer|JSON|
|rebind_timer|JSON|
|push_default_route|Bool|
|push_dns_server|Bool|
|dns_servers_override|StringArray|
|dns_search|StringArray|
|dns_local_name|String|
|zone|String|