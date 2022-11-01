# Table: aws_elbv1_load_balancers

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| created_time | timestamp | X | √ |  | 
| listener_descriptions | json | X | √ |  | 
| security_groups | string_array | X | √ |  | 
| account_id | string | X | √ |  | 
| backend_server_descriptions | json | X | √ |  | 
| availability_zones | string_array | X | √ |  | 
| region | string | X | √ |  | 
| arn | string | √ | √ |  | 
| policies | json | X | √ |  | 
| subnets | string_array | X | √ |  | 
| vpc_id | string | X | √ |  | 
| attributes | json | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| health_check | json | X | √ |  | 
| load_balancer_name | string | X | √ |  | 
| dns_name | string | X | √ |  | 
| instances | json | X | √ |  | 
| scheme | string | X | √ |  | 
| source_security_group | json | X | √ |  | 
| tags | json | X | √ |  | 
| canonical_hosted_zone_name | string | X | √ |  | 
| canonical_hosted_zone_name_id | string | X | √ |  | 


