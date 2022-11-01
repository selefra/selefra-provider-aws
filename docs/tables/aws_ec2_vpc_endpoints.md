# Table: aws_ec2_vpc_endpoints

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| account_id | string | X | √ |  | 
| ip_address_type | string | X | √ |  | 
| policy_document | string | X | √ |  | 
| tags | json | X | √ |  | 
| vpc_endpoint_type | string | X | √ |  | 
| vpc_id | string | X | √ |  | 
| creation_timestamp | timestamp | X | √ |  | 
| dns_options | json | X | √ |  | 
| private_dns_enabled | bool | X | √ |  | 
| requester_managed | bool | X | √ |  | 
| route_table_ids | string_array | X | √ |  | 
| subnet_ids | string_array | X | √ |  | 
| vpc_endpoint_id | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| region | string | X | √ |  | 
| dns_entries | json | X | √ |  | 
| groups | json | X | √ |  | 
| last_error | json | X | √ |  | 
| network_interface_ids | string_array | X | √ |  | 
| service_name | string | X | √ |  | 
| arn | string | √ | √ |  | 
| owner_id | string | X | √ |  | 
| state | string | X | √ |  | 


