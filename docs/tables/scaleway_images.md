# Table: scaleway_images

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
|arch|String|
|creation_date|Timestamp|
|modification_date|Timestamp|
|default_bootscript|JSON|
|extra_volumes|JSON|
|from_server|String|
|organization|String|
|public|Bool|
|root_volume|JSON|
|state|String|
|project|String|
|tags|StringArray|
|zone|String|