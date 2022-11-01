# Table: aws_ec2_network_acls

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| account_id | string | X | √ |  | 
| arn | string | √ | √ |  | 
| associations | json | X | √ |  | 
| network_acl_id | string | X | √ |  | 
| owner_id | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| region | string | X | √ |  | 
| tags | json | X | √ |  | 
| entries | json | X | √ |  | 
| is_default | bool | X | √ |  | 
| vpc_id | string | X | √ |  | 


