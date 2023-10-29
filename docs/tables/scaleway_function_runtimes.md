# Table: scaleway_function_runtimes

This table shows data for Scaleway Function Runtimes.

The composite primary key for this table is (**region**, **name**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|region (PK)|`utf8`|
|name (PK)|`utf8`|
|language|`utf8`|
|version|`utf8`|
|default_handler|`utf8`|
|code_sample|`utf8`|
|status|`utf8`|
|status_message|`utf8`|
|extension|`utf8`|
|implementation|`utf8`|