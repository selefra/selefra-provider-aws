# Table: aws_ec2_security_groups

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| arn | string | √ | √ |  | 
| ip_permissions | json | X | √ |  | 
| owner_id | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| account_id | string | X | √ |  | 
| description | string | X | √ |  | 
| group_id | string | X | √ |  | 
| group_name | string | X | √ |  | 
| ip_permissions_egress | json | X | √ |  | 
| tags | json | X | √ |  | 
| vpc_id | string | X | √ |  | 
| region | string | X | √ |  | 


