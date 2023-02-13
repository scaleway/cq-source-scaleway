# Table: scaleway_function_runtimes

The composite primary key for this table is (**org_id**, **region**, **name**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|org_id (PK)|String|
|region (PK)|String|
|name (PK)|String|
|language|String|
|version|String|
|default_handler|String|
|code_sample|String|
|status|String|
|status_message|String|
|extension|String|
|implementation|String|