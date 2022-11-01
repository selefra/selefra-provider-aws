# Table: aws_glue_ml_transform_task_runs

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| region | string | X | √ |  | 
| last_modified_on | timestamp | X | √ |  | 
| selefra_id | string | √ | √ | random id | 
| aws_glue_ml_transforms_selefra_id | string | X | X | fk to aws_glue_ml_transforms.selefra_id | 
| ml_transform_arn | string | X | √ |  | 
| error_string | string | X | √ |  | 
| execution_time | int | X | √ |  | 
| log_group_name | string | X | √ |  | 
| properties | json | X | √ |  | 
| account_id | string | X | √ |  | 
| completed_on | timestamp | X | √ |  | 
| started_on | timestamp | X | √ |  | 
| status | string | X | √ |  | 
| task_run_id | string | X | √ |  | 
| transform_id | string | X | √ |  | 


