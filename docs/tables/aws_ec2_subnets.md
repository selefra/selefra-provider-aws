# Table: aws_ec2_subnets

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| arn | string | √ | √ |  | 
| assign_ipv6_address_on_creation | bool | X | √ |  | 
| ipv6_cidr_block_association_set | json | X | √ |  | 
| ipv6_native | bool | X | √ |  | 
| map_public_ip_on_launch | bool | X | √ |  | 
| tags | json | X | √ |  | 
| account_id | string | X | √ |  | 
| cidr_block | string | X | √ |  | 
| customer_owned_ipv4_pool | string | X | √ |  | 
| default_for_az | bool | X | √ |  | 
| enable_dns64 | bool | X | √ |  | 
| state | string | X | √ |  | 
| subnet_arn | string | X | √ |  | 
| region | string | X | √ |  | 
| map_customer_owned_ip_on_launch | bool | X | √ |  | 
| owner_id | string | X | √ |  | 
| private_dns_name_options_on_launch | json | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| availability_zone | string | X | √ |  | 
| availability_zone_id | string | X | √ |  | 
| available_ip_address_count | int | X | √ |  | 
| enable_lni_at_device_index | int | X | √ |  | 
| outpost_arn | string | X | √ |  | 
| subnet_id | string | X | √ |  | 
| vpc_id | string | X | √ |  | 


