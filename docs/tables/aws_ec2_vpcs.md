# Table: aws_ec2_vpcs

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| instance_tenancy | string | X | √ |  | 
| owner_id | string | X | √ |  | 
| tags | json | X | √ |  | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| arn | string | √ | √ |  | 
| ipv6_cidr_block_association_set | json | X | √ |  | 
| cidr_block_association_set | json | X | √ |  | 
| state | string | X | √ |  | 
| vpc_id | string | X | √ |  | 
| cidr_block | string | X | √ |  | 
| dhcp_options_id | string | X | √ |  | 
| is_default | bool | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 


