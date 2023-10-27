# Table: scaleway_iot_networks

This table shows data for Scaleway IOT Networks.

The primary key for this table is **id**.

## Relations

This table depends on [scaleway_iot_hubs](scaleway_iot_hubs.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|id (PK)|`utf8`|
|name|`utf8`|
|type|`utf8`|
|endpoint|`utf8`|
|hub_id|`utf8`|
|created_at|`timestamp[us, tz=UTC]`|
|topic_prefix|`utf8`|