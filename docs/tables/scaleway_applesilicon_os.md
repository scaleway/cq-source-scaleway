# Table: scaleway_applesilicon_os

This table shows data for Scaleway Applesilicon Os.

The composite primary key for this table is (**zone**, **id**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|zone (PK)|`utf8`|
|id (PK)|`utf8`|
|name|`utf8`|
|label|`utf8`|
|image_url|`utf8`|
|compatible_server_types|`list<item: utf8, nullable>`|