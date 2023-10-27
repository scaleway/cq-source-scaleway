# Table: scaleway_instances

This table shows data for Scaleway Instances.

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|id (PK)|`utf8`|
|name|`utf8`|
|organization|`utf8`|
|project|`utf8`|
|allowed_actions|`list<item: utf8, nullable>`|
|tags|`list<item: utf8, nullable>`|
|commercial_type|`utf8`|
|creation_date|`timestamp[us, tz=UTC]`|
|dynamic_ip_required|`bool`|
|enable_ipv6|`bool`|
|hostname|`utf8`|
|image|`json`|
|protected|`bool`|
|private_ip|`utf8`|
|public_ip|`json`|
|modification_date|`timestamp[us, tz=UTC]`|
|state|`utf8`|
|location|`json`|
|ipv6|`json`|
|bootscript|`json`|
|boot_type|`utf8`|
|volumes|`json`|
|security_group|`json`|
|maintenances|`json`|
|state_detail|`utf8`|
|arch|`utf8`|
|placement_group|`json`|
|private_nics|`json`|
|zone|`utf8`|