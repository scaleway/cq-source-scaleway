# Table: scaleway_baremetal_os

This table shows data for Scaleway Baremetal Os.

The composite primary key for this table is (**zone**, **id**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|zone (PK)|`utf8`|
|id (PK)|`utf8`|
|name|`utf8`|
|version|`utf8`|
|logo_url|`utf8`|
|ssh|`json`|
|user|`json`|
|password|`json`|
|service_user|`json`|
|service_password|`json`|
|enabled|`bool`|
|license_required|`bool`|