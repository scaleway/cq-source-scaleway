# Table: scaleway_lb_certificates

This table shows data for Scaleway Lb Certificates.

The composite primary key for this table is (**lb_id**, **id**).

## Relations

This table depends on [scaleway_lbs](scaleway_lbs.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|lb_id (PK)|`utf8`|
|type|`utf8`|
|id (PK)|`utf8`|
|common_name|`utf8`|
|subject_alternative_name|`list<item: utf8, nullable>`|
|fingerprint|`utf8`|
|not_valid_before|`timestamp[us, tz=UTC]`|
|not_valid_after|`timestamp[us, tz=UTC]`|
|status|`utf8`|
|lb|`json`|
|name|`utf8`|
|created_at|`timestamp[us, tz=UTC]`|
|updated_at|`timestamp[us, tz=UTC]`|
|status_details|`utf8`|