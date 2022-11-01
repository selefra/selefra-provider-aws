# Table: aws_cloudtrail_trails

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| s3_key_prefix | string | X | √ |  | 
| sns_topic_name | string | X | √ |  | 
| tags | json | X | √ |  | 
| include_global_service_events | bool | X | √ |  | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| cloudwatch_logs_log_group_name | string | X | √ |  | 
| status | json | X | √ |  | 
| cloud_watch_logs_role_arn | string | X | √ |  | 
| has_custom_event_selectors | bool | X | √ |  | 
| has_insight_selectors | bool | X | √ |  | 
| log_file_validation_enabled | bool | X | √ |  | 
| s3_bucket_name | string | X | √ |  | 
| cloud_watch_logs_log_group_arn | string | X | √ |  | 
| home_region | string | X | √ |  | 
| is_organization_trail | bool | X | √ |  | 
| sns_topic_arn | string | X | √ |  | 
| arn | string | √ | √ |  | 
| is_multi_region_trail | bool | X | √ |  | 
| kms_key_id | string | X | √ |  | 
| name | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 


