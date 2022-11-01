# Table: aws_elasticache_service_updates

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| service_update_severity | string | X | √ |  | 
| region | string | X | √ |  | 
| service_update_name | string | X | √ |  | 
| service_update_description | string | X | √ |  | 
| service_update_status | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| engine_version | string | X | √ |  | 
| estimated_update_time | string | X | √ |  | 
| service_update_type | string | X | √ |  | 
| arn | string | √ | √ |  | 
| service_update_end_date | timestamp | X | √ |  | 
| engine | string | X | √ |  | 
| service_update_recommended_apply_by_date | timestamp | X | √ |  | 
| service_update_release_date | timestamp | X | √ |  | 
| account_id | string | X | √ |  | 
| auto_update_after_recommended_apply_by_date | bool | X | √ |  | 


