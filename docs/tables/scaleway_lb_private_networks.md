# Table: scaleway_lb_private_networks

This table shows data for Scaleway Lb Private Networks.

The composite primary key for this table is (**lb_id**, **private_network_id**).

## Relations

This table depends on [scaleway_lbs](scaleway_lbs.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|lb_id (PK)|`utf8`|
|lb|`json`|
|static_config|`json`|
|dhcp_config|`json`|
|private_network_id (PK)|`utf8`|
|status|`utf8`|
|created_at|`timestamp[us, tz=UTC]`|
|updated_at|`timestamp[us, tz=UTC]`|