# Table: aws_elasticache_user_groups

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| account_id | string | X | √ |  | 
| arn | string | √ | √ |  | 
| replication_groups | string_array | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| region | string | X | √ |  | 
| engine | string | X | √ |  | 
| minimum_engine_version | string | X | √ |  | 
| pending_changes | json | X | √ |  | 
| status | string | X | √ |  | 
| user_group_id | string | X | √ |  | 
| user_ids | string_array | X | √ |  | 


