# Table: aws_rds_subnet_groups

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| account_id | string | X | √ |  | 
| db_subnet_group_description | string | X | √ |  | 
| db_subnet_group_name | string | X | √ |  | 
| subnet_group_status | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| region | string | X | √ |  | 
| arn | string | √ | √ |  | 
| subnets | json | X | √ |  | 
| supported_network_types | string_array | X | √ |  | 
| vpc_id | string | X | √ |  | 


