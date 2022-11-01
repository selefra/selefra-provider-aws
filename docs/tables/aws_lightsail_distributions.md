# Table: aws_lightsail_distributions

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| account_id | string | X | √ |  | 
| tags | json | X | √ |  | 
| bundle_id | string | X | √ |  | 
| default_cache_behavior | json | X | √ |  | 
| ip_address_type | string | X | √ |  | 
| resource_type | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| region | string | X | √ |  | 
| able_to_update_bundle | bool | X | √ |  | 
| arn | string | √ | √ |  | 
| alternative_domain_names | string_array | X | √ |  | 
| cache_behavior_settings | json | X | √ |  | 
| domain_name | string | X | √ |  | 
| location | json | X | √ |  | 
| origin | json | X | √ |  | 
| cache_behaviors | json | X | √ |  | 
| certificate_name | string | X | √ |  | 
| created_at | timestamp | X | √ |  | 
| is_enabled | bool | X | √ |  | 
| name | string | X | √ |  | 
| origin_public_dns | string | X | √ |  | 
| status | string | X | √ |  | 
| support_code | string | X | √ |  | 
| latest_cache_reset | json | X | √ |  | 


