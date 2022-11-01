# Table: aws_athena_work_group_named_queries

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| work_group_arn | string | X | √ |  | 
| named_query_id | string | X | √ |  | 
| selefra_id | string | √ | √ | random id | 
| aws_athena_work_groups_selefra_id | string | X | X | fk to aws_athena_work_groups.selefra_id | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| database | string | X | √ |  | 
| name | string | X | √ |  | 
| query_string | string | X | √ |  | 
| description | string | X | √ |  | 
| work_group | string | X | √ |  | 


