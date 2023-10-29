# Table: scaleway_applesilicon_server_types

This table shows data for Scaleway Applesilicon Server Types.

The composite primary key for this table is (**zone**, **name**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|zone (PK)|`utf8`|
|cpu|`json`|
|disk|`json`|
|name (PK)|`utf8`|
|memory|`json`|
|stock|`utf8`|
|minimum_lease_duration|`json`|