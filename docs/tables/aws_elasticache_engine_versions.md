# Table: aws_elasticache_engine_versions

## Primary Keys 

```
account_id, region
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| cache_parameter_group_family | string | X | √ |  | 
| engine | string | X | √ |  | 
| engine_version | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| account_id | string | X | √ | `The AWS Account ID of the resource.` | 
| region | string | X | √ | `The AWS Region of the resource.` | 
| cache_engine_description | string | X | √ |  | 
| cache_engine_version_description | string | X | √ |  | 


