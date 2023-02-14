# Table: scaleway_instances

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|id (PK)|String|
|name|String|
|organization|String|
|project|String|
|allowed_actions|StringArray|
|tags|StringArray|
|commercial_type|String|
|creation_date|Timestamp|
|dynamic_ip_required|Bool|
|enable_ipv6|Bool|
|hostname|String|
|image|JSON|
|protected|Bool|
|private_ip|String|
|public_ip|JSON|
|modification_date|Timestamp|
|state|String|
|location|JSON|
|ipv6|JSON|
|bootscript|JSON|
|boot_type|String|
|volumes|JSON|
|security_group|JSON|
|maintenances|JSON|
|state_detail|String|
|arch|String|
|placement_group|JSON|
|private_nics|JSON|
|zone|String|