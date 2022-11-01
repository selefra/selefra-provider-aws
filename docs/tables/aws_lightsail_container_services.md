# Table: aws_lightsail_container_services

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| principal_arn | string | X | √ |  | 
| private_domain_name | string | X | √ |  | 
| scale | int | X | √ |  | 
| state | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| location | json | X | √ |  | 
| public_domain_names | json | X | √ |  | 
| created_at | timestamp | X | √ |  | 
| power | string | X | √ |  | 
| private_registry_access | json | X | √ |  | 
| resource_type | string | X | √ |  | 
| url | string | X | √ |  | 
| arn | string | √ | √ |  | 
| tags | json | X | √ |  | 
| next_deployment | json | X | √ |  | 
| power_id | string | X | √ |  | 
| state_detail | json | X | √ |  | 
| container_service_name | string | X | √ |  | 
| current_deployment | json | X | √ |  | 
| is_disabled | bool | X | √ |  | 


