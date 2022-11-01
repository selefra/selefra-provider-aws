# Table: aws_dynamodb_table_continuous_backups

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| continuous_backups_status | string | X | √ |  | 
| point_in_time_recovery_description | json | X | √ |  | 
| selefra_id | string | √ | √ | random id | 
| aws_dynamodb_tables_selefra_id | string | X | X | fk to aws_dynamodb_tables.selefra_id | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| table_arn | string | X | √ |  | 


