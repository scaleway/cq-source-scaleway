# Table: scaleway_tem_emails

The primary key for this table is **id**.
It supports incremental syncs.
## Relations

This table depends on [scaleway_tem_domains](scaleway_tem_domains.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|id (PK)|String|
|message_id|String|
|project_id|String|
|mail_from|String|
|rcpt_to|String|
|rcpt_type|String|
|created_at|Timestamp|
|updated_at|Timestamp|
|status|String|
|status_details|String|
|try_count|Int|
|last_tries|JSON|