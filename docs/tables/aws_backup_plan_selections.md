# Table: aws_backup_plan_selections

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| region | string | X | √ |  | 
| plan_arn | string | X | √ |  | 
| backup_selection | json | X | √ |  | 
| creation_date | timestamp | X | √ |  | 
| selection_id | string | X | √ |  | 
| aws_backup_plans_selefra_id | string | X | X | fk to aws_backup_plans.selefra_id | 
| account_id | string | X | √ |  | 
| backup_plan_id | string | X | √ |  | 
| creator_request_id | string | X | √ |  | 
| result_metadata | json | X | √ |  | 
| selefra_id | string | √ | √ | random id | 


