# Table: aws_ec2_vpc_peering_connections

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| requester_vpc_info | json | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| arn | string | √ | √ |  | 
| accepter_vpc_info | json | X | √ |  | 
| expiration_time | timestamp | X | √ |  | 
| status | json | X | √ |  | 
| tags | json | X | √ |  | 
| vpc_peering_connection_id | string | X | √ |  | 


