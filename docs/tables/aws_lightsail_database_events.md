# Table: aws_lightsail_database_events

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| region | string | X | √ |  | 
| message | string | X | √ |  | 
| aws_lightsail_databases_selefra_id | string | X | X | fk to aws_lightsail_databases.selefra_id | 
| resource | string | X | √ |  | 
| selefra_id | string | √ | √ | random id | 
| account_id | string | X | √ |  | 
| database_arn | string | X | √ |  | 
| created_at | timestamp | X | √ |  | 
| event_categories | string_array | X | √ |  | 


