# Table: scaleway_tem_emails

This table shows data for Scaleway Tem Emails.

The primary key for this table is **id**.
It supports incremental syncs.
## Relations

This table depends on [scaleway_tem_domains](scaleway_tem_domains.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|id (PK)|`utf8`|
|message_id|`utf8`|
|project_id|`utf8`|
|mail_from|`utf8`|
|rcpt_to|`utf8`|
|rcpt_type|`utf8`|
|created_at|`timestamp[us, tz=UTC]`|
|updated_at|`timestamp[us, tz=UTC]`|
|status|`utf8`|
|status_details|`utf8`|
|try_count|`int64`|
|last_tries|`json`|