# Table: aws_elbv2_load_balancers

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| account_id | string | X | √ |  | 
| web_acl_arn | string | X | √ |  | 
| customer_owned_ipv4_pool | string | X | √ |  | 
| state | json | X | √ |  | 
| tags | json | X | √ |  | 
| arn | string | √ | √ |  | 
| ip_address_type | string | X | √ |  | 
| security_groups | string_array | X | √ |  | 
| region | string | X | √ |  | 
| availability_zones | json | X | √ |  | 
| canonical_hosted_zone_id | string | X | √ |  | 
| created_time | timestamp | X | √ |  | 
| scheme | string | X | √ |  | 
| dns_name | string | X | √ |  | 
| load_balancer_name | string | X | √ |  | 
| type | string | X | √ |  | 
| vpc_id | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 


