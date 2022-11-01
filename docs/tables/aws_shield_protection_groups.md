# Table: aws_shield_protection_groups

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| account_id | string | X | √ |  | 
| arn | string | √ | √ |  | 
| tags | json | X | √ |  | 
| aggregation | string | X | √ |  | 
| members | string_array | X | √ |  | 
| resource_type | string | X | √ |  | 
| pattern | string | X | √ |  | 
| protection_group_id | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 


