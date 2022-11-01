# Table: aws_ec2_nat_gateways

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| nat_gateway_addresses | json | X | √ |  | 
| provisioned_bandwidth | json | X | √ |  | 
| arn | string | √ | √ |  | 
| failure_code | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| delete_time | timestamp | X | √ |  | 
| failure_message | string | X | √ |  | 
| vpc_id | string | X | √ |  | 
| connectivity_type | string | X | √ |  | 
| create_time | timestamp | X | √ |  | 
| nat_gateway_id | string | X | √ |  | 
| state | string | X | √ |  | 
| subnet_id | string | X | √ |  | 
| tags | json | X | √ |  | 


