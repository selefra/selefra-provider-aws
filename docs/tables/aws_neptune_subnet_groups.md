# Table: aws_neptune_subnet_groups

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| account_id | string | X | √ |  | 
| arn | string | √ | √ |  | 
| db_subnet_group_description | string | X | √ |  | 
| db_subnet_group_name | string | X | √ |  | 
| vpc_id | string | X | √ |  | 
| region | string | X | √ |  | 
| tags | json | X | √ |  | 
| subnet_group_status | string | X | √ |  | 
| subnets | json | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 


