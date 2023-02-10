# Table: scaleway_vpc_private_networks

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
|organization_id|String|
|project_id|String|
|zone|String|
|tags|StringArray|
|created_at|Timestamp|
|updated_at|Timestamp|
|subnets|JSON|