# Table: aws_route53_domains

## Primary Keys 

```
account_id, domain_name
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| admin_contact | json | X | √ |  | 
| tech_contact | json | X | √ |  | 
| abuse_contact_email | string | X | √ |  | 
| expiration_date | timestamp | X | √ |  | 
| registrar_name | string | X | √ |  | 
| updated_date | timestamp | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| account_id | string | X | √ |  | 
| admin_privacy | bool | X | √ |  | 
| creation_date | timestamp | X | √ |  | 
| registrant_privacy | bool | X | √ |  | 
| registry_domain_id | string | X | √ |  | 
| reseller | string | X | √ |  | 
| registrant_contact | json | X | √ |  | 
| dns_sec | string | X | √ |  | 
| registrar_url | string | X | √ |  | 
| status_list | string_array | X | √ |  | 
| who_is_server | string | X | √ |  | 
| result_metadata | json | X | √ |  | 
| tags | json | X | √ | `A list of tags` | 
| nameservers | json | X | √ |  | 
| abuse_contact_phone | string | X | √ |  | 
| auto_renew | bool | X | √ |  | 
| tech_privacy | bool | X | √ |  | 
| domain_name | string | X | √ |  | 


