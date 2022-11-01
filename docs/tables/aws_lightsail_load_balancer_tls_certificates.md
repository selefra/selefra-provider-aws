# Table: aws_lightsail_load_balancer_tls_certificates

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| issuer | string | X | √ |  | 
| load_balancer_name | string | X | √ |  | 
| arn | string | X | √ |  | 
| failure_reason | string | X | √ |  | 
| revocation_reason | string | X | √ |  | 
| serial | string | X | √ |  | 
| domain_validation_records | json | X | √ |  | 
| issued_at | timestamp | X | √ |  | 
| location | json | X | √ |  | 
| load_balancer_arn | string | X | √ |  | 
| tags | json | X | √ |  | 
| is_attached | bool | X | √ |  | 
| key_algorithm | string | X | √ |  | 
| not_before | timestamp | X | √ |  | 
| revoked_at | timestamp | X | √ |  | 
| support_code | string | X | √ |  | 
| created_at | timestamp | X | √ |  | 
| domain_name | string | X | √ |  | 
| name | string | X | √ |  | 
| not_after | timestamp | X | √ |  | 
| resource_type | string | X | √ |  | 
| selefra_id | string | √ | √ | random id | 
| subject_alternative_names | string_array | X | √ |  | 
| aws_lightsail_load_balancers_selefra_id | string | X | X | fk to aws_lightsail_load_balancers.selefra_id | 
| renewal_summary | json | X | √ |  | 
| signature_algorithm | string | X | √ |  | 
| status | string | X | √ |  | 
| subject | string | X | √ |  | 


