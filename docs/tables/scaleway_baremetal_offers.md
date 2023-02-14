# Table: scaleway_baremetal_offers

The composite primary key for this table is (**zone**, **id**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|zone (PK)|String|
|id (PK)|String|
|name|String|
|stock|String|
|bandwidth|Int|
|commercial_range|String|
|price_per_hour|JSON|
|price_per_month|JSON|
|disks|JSON|
|enable|Bool|
|cpus|JSON|
|memories|JSON|
|quota_name|String|
|persistent_memories|JSON|
|raid_controllers|JSON|
|incompatible_os_ids|StringArray|
|subscription_period|String|
|operation_path|String|
|fee|JSON|
|options|JSON|