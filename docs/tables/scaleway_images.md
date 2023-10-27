# Table: scaleway_images

This table shows data for Scaleway Images.

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|id (PK)|`utf8`|
|name|`utf8`|
|arch|`utf8`|
|creation_date|`timestamp[us, tz=UTC]`|
|modification_date|`timestamp[us, tz=UTC]`|
|default_bootscript|`json`|
|extra_volumes|`json`|
|from_server|`utf8`|
|organization|`utf8`|
|public|`bool`|
|root_volume|`json`|
|state|`utf8`|
|project|`utf8`|
|tags|`list<item: utf8, nullable>`|
|zone|`utf8`|