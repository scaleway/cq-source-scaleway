# Table: scaleway_container_crons

This table shows data for Scaleway Container Crons.

The primary key for this table is **id**.

## Relations

This table depends on [scaleway_containers](scaleway_containers.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|id (PK)|`utf8`|
|container_id|`utf8`|
|schedule|`utf8`|
|args|`json`|
|status|`utf8`|
|name|`utf8`|