# Table: scaleway_tem_domains

The primary key for this table is **id**.

## Relations

The following tables depend on scaleway_tem_domains:
  - [scaleway_tem_emails](scaleway_tem_emails.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|id (PK)|String|
|organization_id|String|
|project_id|String|
|name|String|
|status|String|
|created_at|Timestamp|
|next_check_at|Timestamp|
|last_valid_at|Timestamp|
|revoked_at|Timestamp|
|last_error|String|
|spf_config|String|
|dkim_config|String|
|statistics|JSON|
|region|String|