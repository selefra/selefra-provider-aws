# Table: aws_athena_work_group_prepared_statements

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| region | string | X | √ |  | 
| description | string | X | √ |  | 
| last_modified_time | timestamp | X | √ |  | 
| work_group_name | string | X | √ |  | 
| aws_athena_work_groups_selefra_id | string | X | X | fk to aws_athena_work_groups.selefra_id | 
| account_id | string | X | √ |  | 
| work_group_arn | string | X | √ |  | 
| query_statement | string | X | √ |  | 
| statement_name | string | X | √ |  | 
| selefra_id | string | √ | √ | random id | 


