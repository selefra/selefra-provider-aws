# Table: aws_lightsail_instances

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| tags | json | X | √ |  | 
| location | json | X | √ |  | 
| metadata_options | json | X | √ |  | 
| name | string | X | √ |  | 
| access_details | json | X | √ |  | 
| hardware | json | X | √ |  | 
| resource_type | string | X | √ |  | 
| support_code | string | X | √ |  | 
| region | string | X | √ |  | 
| blueprint_id | string | X | √ |  | 
| created_at | timestamp | X | √ |  | 
| ipv6_addresses | string_array | X | √ |  | 
| is_static_ip | bool | X | √ |  | 
| networking | json | X | √ |  | 
| state | json | X | √ |  | 
| public_ip_address | string | X | √ |  | 
| account_id | string | X | √ |  | 
| arn | string | √ | √ |  | 
| add_ons | json | X | √ |  | 
| blueprint_name | string | X | √ |  | 
| bundle_id | string | X | √ |  | 
| ip_address_type | string | X | √ |  | 
| private_ip_address | string | X | √ |  | 
| ssh_key_name | string | X | √ |  | 
| username | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 


