# Table: scaleway_mnq_namespace_credentials

The composite primary key for this table is (**id**, **namespace_id**).

## Relations

This table depends on [scaleway_mnq_namespaces](scaleway_mnq_namespaces.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|id (PK)|String|
|name|String|
|namespace_id (PK)|String|
|protocol|String|
|nats_credentials|JSON|
|sqs_sns_credentials|JSON|