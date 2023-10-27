# Table: scaleway_lb_frontend_acls

This table shows data for Scaleway Lb Frontend ACLs.

The composite primary key for this table is (**frontend_id**, **id**).

## Relations

This table depends on [scaleway_lb_frontends](scaleway_lb_frontends.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|frontend_id (PK)|`utf8`|
|id (PK)|`utf8`|
|name|`utf8`|
|match|`json`|
|action|`json`|
|frontend|`json`|
|index|`int64`|
|created_at|`timestamp[us, tz=UTC]`|
|updated_at|`timestamp[us, tz=UTC]`|
|description|`utf8`|