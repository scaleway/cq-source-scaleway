# Table: scaleway_baremetal_os

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
|version|String|
|logo_url|String|
|ssh|JSON|
|user|JSON|
|password|JSON|
|service_user|JSON|
|service_password|JSON|
|enabled|Bool|
|license_required|Bool|