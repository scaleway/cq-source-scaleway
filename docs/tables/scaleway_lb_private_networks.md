# Table: scaleway_lb_private_networks

The composite primary key for this table is (**lb_id**, **private_network_id**).

## Relations

This table depends on [scaleway_lbs](scaleway_lbs.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|lb_id (PK)|String|
|lb|JSON|
|static_config|JSON|
|dhcp_config|JSON|
|private_network_id (PK)|String|
|status|String|
|created_at|Timestamp|
|updated_at|Timestamp|