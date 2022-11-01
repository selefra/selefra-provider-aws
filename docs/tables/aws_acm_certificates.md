# Table: aws_acm_certificates

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| renewal_summary | json | X | √ |  | 
| signature_algorithm | string | X | √ |  | 
| subject | string | X | √ |  | 
| tags | json | X | √ |  | 
| renewal_eligibility | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| not_after | timestamp | X | √ |  | 
| revoked_at | timestamp | X | √ |  | 
| subject_alternative_names | string_array | X | √ |  | 
| domain_name | string | X | √ |  | 
| extended_key_usages | json | X | √ |  | 
| issuer | string | X | √ |  | 
| not_before | timestamp | X | √ |  | 
| account_id | string | X | √ |  | 
| created_at | timestamp | X | √ |  | 
| domain_validation_options | json | X | √ |  | 
| imported_at | timestamp | X | √ |  | 
| options | json | X | √ |  | 
| serial | string | X | √ |  | 
| status | string | X | √ |  | 
| arn | string | √ | √ |  | 
| in_use_by | string_array | X | √ |  | 
| revocation_reason | string | X | √ |  | 
| region | string | X | √ |  | 
| certificate_authority_arn | string | X | √ |  | 
| failure_reason | string | X | √ |  | 
| issued_at | timestamp | X | √ |  | 
| key_algorithm | string | X | √ |  | 
| key_usages | json | X | √ |  | 
| type | string | X | √ |  | 


