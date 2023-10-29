# Table: scaleway_tem_domains

This table shows data for Scaleway Tem Domains.

The primary key for this table is **id**.

## Relations

The following tables depend on scaleway_tem_domains:
  - [scaleway_tem_emails](scaleway_tem_emails.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|id (PK)|`utf8`|
|organization_id|`utf8`|
|project_id|`utf8`|
|name|`utf8`|
|status|`utf8`|
|created_at|`timestamp[us, tz=UTC]`|
|next_check_at|`timestamp[us, tz=UTC]`|
|last_valid_at|`timestamp[us, tz=UTC]`|
|revoked_at|`timestamp[us, tz=UTC]`|
|last_error|`utf8`|
|spf_config|`utf8`|
|dkim_config|`utf8`|
|statistics|`json`|
|region|`utf8`|