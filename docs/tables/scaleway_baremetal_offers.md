# Table: scaleway_baremetal_offers

This table shows data for Scaleway Baremetal Offers.

The composite primary key for this table is (**zone**, **id**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|zone (PK)|`utf8`|
|id (PK)|`utf8`|
|name|`utf8`|
|stock|`utf8`|
|bandwidth|`int64`|
|commercial_range|`utf8`|
|price_per_hour|`json`|
|price_per_month|`json`|
|disks|`json`|
|enable|`bool`|
|cpus|`json`|
|memories|`json`|
|quota_name|`utf8`|
|persistent_memories|`json`|
|raid_controllers|`json`|
|incompatible_os_ids|`list<item: utf8, nullable>`|
|subscription_period|`utf8`|
|operation_path|`utf8`|
|fee|`json`|
|options|`json`|