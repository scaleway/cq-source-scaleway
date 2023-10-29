# Table: scaleway_mnq_namespace_credentials

This table shows data for Scaleway Mnq Namespace Credentials.

The composite primary key for this table is (**id**, **namespace_id**).

## Relations

This table depends on [scaleway_mnq_namespaces](scaleway_mnq_namespaces.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|id (PK)|`utf8`|
|name|`utf8`|
|namespace_id (PK)|`utf8`|
|protocol|`utf8`|
|nats_credentials|`json`|
|sqs_sns_credentials|`json`|