# Table: aws_lightsail_certificates

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| account_id | string | X | √ |  | 
| arn | string | X | √ |  | 
| domain_name | string | X | √ |  | 
| not_before | timestamp | X | √ |  | 
| domain_validation_records | json | X | √ |  | 
| eligible_to_renew | string | X | √ |  | 
| issued_at | timestamp | X | √ |  | 
| not_after | timestamp | X | √ |  | 
| request_failure_reason | string | X | √ |  | 
| in_use_resource_count | int | X | √ |  | 
| issuer_ca | string | X | √ |  | 
| key_algorithm | string | X | √ |  | 
| renewal_summary | json | X | √ |  | 
| revocation_reason | string | X | √ |  | 
| revoked_at | timestamp | X | √ |  | 
| selefra_id | string | √ | √ | random id | 
| support_code | string | X | √ |  | 
| region | string | X | √ |  | 
| tags | json | X | √ |  | 
| created_at | timestamp | X | √ |  | 
| name | string | X | √ |  | 
| serial_number | string | X | √ |  | 
| status | string | X | √ |  | 
| subject_alternative_names | string_array | X | √ |  | 


