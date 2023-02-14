# Table: scaleway_lb_certificates

The composite primary key for this table is (**lb_id**, **id**).

## Relations

This table depends on [scaleway_lbs](scaleway_lbs.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|lb_id (PK)|String|
|type|String|
|id (PK)|String|
|common_name|String|
|subject_alternative_name|StringArray|
|fingerprint|String|
|not_valid_before|Timestamp|
|not_valid_after|Timestamp|
|status|String|
|lb|JSON|
|name|String|
|created_at|Timestamp|
|updated_at|Timestamp|
|status_details|String|