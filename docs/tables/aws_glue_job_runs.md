# Table: aws_glue_job_runs

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| execution_class | string | X | √ |  | 
| job_run_state | string | X | √ |  | 
| arguments | json | X | √ |  | 
| previous_run_id | string | X | √ |  | 
| aws_glue_jobs_selefra_id | string | X | X | fk to aws_glue_jobs.selefra_id | 
| glue_version | string | X | √ |  | 
| id | string | X | √ |  | 
| number_of_workers | int | X | √ |  | 
| started_on | timestamp | X | √ |  | 
| worker_type | string | X | √ |  | 
| account_id | string | X | √ |  | 
| job_arn | string | X | √ |  | 
| allocated_capacity | int | X | √ |  | 
| dpu_seconds | float | X | √ |  | 
| job_name | string | X | √ |  | 
| notification_property | json | X | √ |  | 
| attempt | int | X | √ |  | 
| log_group_name | string | X | √ |  | 
| security_configuration | string | X | √ |  | 
| selefra_id | string | √ | √ | random id | 
| region | string | X | √ |  | 
| predecessor_runs | json | X | √ |  | 
| error_message | string | X | √ |  | 
| execution_time | int | X | √ |  | 
| max_capacity | float | X | √ |  | 
| trigger_name | string | X | √ |  | 
| completed_on | timestamp | X | √ |  | 
| last_modified_on | timestamp | X | √ |  | 
| timeout | int | X | √ |  | 


