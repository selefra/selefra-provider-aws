# Table: aws_cloudwatchlogs_log_groups

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| kms_key_id | string | X | √ |  | 
| log_group_name | string | X | √ |  | 
| metric_filter_count | int | X | √ |  | 
| retention_in_days | int | X | √ |  | 
| account_id | string | X | √ |  | 
| arn | string | √ | √ |  | 
| tags | json | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| region | string | X | √ |  | 
| creation_time | int | X | √ |  | 
| stored_bytes | int | X | √ |  | 


