# Table: aws_ec2_network_interfaces

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| subnet_id | string | X | √ |  | 
| description | string | X | √ |  | 
| ipv6_addresses | json | X | √ |  | 
| outpost_arn | string | X | √ |  | 
| network_interface_id | string | X | √ |  | 
| ipv6_native | bool | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| tags | json | X | √ |  | 
| association | json | X | √ |  | 
| attachment | json | X | √ |  | 
| groups | json | X | √ |  | 
| ipv4_prefixes | json | X | √ |  | 
| private_dns_name | string | X | √ |  | 
| requester_managed | bool | X | √ |  | 
| status | string | X | √ |  | 
| arn | string | √ | √ |  | 
| vpc_id | string | X | √ |  | 
| interface_type | string | X | √ |  | 
| mac_address | string | X | √ |  | 
| region | string | X | √ |  | 
| ipv6_prefixes | json | X | √ |  | 
| owner_id | string | X | √ |  | 
| private_ip_address | string | X | √ |  | 
| requester_id | string | X | √ |  | 
| deny_all_igw_traffic | bool | X | √ |  | 
| availability_zone | string | X | √ |  | 
| ipv6_address | string | X | √ |  | 
| private_ip_addresses | json | X | √ |  | 
| source_dest_check | bool | X | √ |  | 
| account_id | string | X | √ |  | 


