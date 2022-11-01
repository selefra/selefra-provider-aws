# Table: aws_ec2_vpc_endpoint_services

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| tags | json | X | √ |  | 
| acceptance_required | bool | X | √ |  | 
| owner | string | X | √ |  | 
| vpc_endpoint_policy_supported | bool | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| region | string | X | √ |  | 
| base_endpoint_dns_names | string_array | X | √ |  | 
| private_dns_name | string | X | √ |  | 
| private_dns_name_verification_state | string | X | √ |  | 
| arn | string | √ | √ |  | 
| service_name | string | X | √ |  | 
| account_id | string | X | √ |  | 
| manages_vpc_endpoints | bool | X | √ |  | 
| payer_responsibility | string | X | √ |  | 
| private_dns_names | json | X | √ |  | 
| service_id | string | X | √ |  | 
| service_type | json | X | √ |  | 
| supported_ip_address_types | string_array | X | √ |  | 
| availability_zones | string_array | X | √ |  | 


