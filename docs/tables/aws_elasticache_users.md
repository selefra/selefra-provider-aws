# Table: aws_elasticache_users

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| access_string | string | X | √ |  | 
| authentication | json | X | √ |  | 
| engine | string | X | √ |  | 
| status | string | X | √ |  | 
| user_group_ids | string_array | X | √ |  | 
| user_id | string | X | √ |  | 
| user_name | string | X | √ |  | 
| arn | string | √ | √ |  | 
| minimum_engine_version | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 


