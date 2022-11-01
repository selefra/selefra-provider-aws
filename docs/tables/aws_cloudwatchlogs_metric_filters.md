# Table: aws_cloudwatchlogs_metric_filters

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| filter_name | string | X | √ |  | 
| filter_pattern | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| arn | string | √ | √ |  | 
| creation_time | int | X | √ |  | 
| log_group_name | string | X | √ |  | 
| metric_transformations | json | X | √ |  | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 


