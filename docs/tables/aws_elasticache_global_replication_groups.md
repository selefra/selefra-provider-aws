# Table: aws_elasticache_global_replication_groups

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| auth_token_enabled | bool | X | √ |  | 
| engine_version | string | X | √ |  | 
| members | json | X | √ |  | 
| transit_encryption_enabled | bool | X | √ |  | 
| cache_node_type | string | X | √ |  | 
| global_replication_group_description | string | X | √ |  | 
| status | string | X | √ |  | 
| at_rest_encryption_enabled | bool | X | √ |  | 
| global_node_groups | json | X | √ |  | 
| global_replication_group_id | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| arn | string | √ | √ |  | 
| cluster_enabled | bool | X | √ |  | 
| engine | string | X | √ |  | 


