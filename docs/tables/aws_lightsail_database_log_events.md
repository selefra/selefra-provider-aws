# Table: aws_lightsail_database_log_events

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| region | string | X | √ |  | 
| database_arn | string | X | √ |  | 
| created_at | timestamp | X | √ |  | 
| message | string | X | √ |  | 
| log_stream_name | string | X | √ |  | 
| aws_lightsail_databases_selefra_id | string | X | X | fk to aws_lightsail_databases.selefra_id | 
| selefra_id | string | √ | √ | random id | 
| account_id | string | X | √ |  | 


