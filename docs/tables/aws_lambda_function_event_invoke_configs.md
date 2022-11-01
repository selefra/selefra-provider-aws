# Table: aws_lambda_function_event_invoke_configs

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| account_id | string | X | √ |  | 
| last_modified | timestamp | X | √ |  | 
| maximum_event_age_in_seconds | int | X | √ |  | 
| maximum_retry_attempts | int | X | √ |  | 
| region | string | X | √ |  | 
| function_arn | string | X | √ |  | 
| destination_config | json | X | √ |  | 
| aws_lambda_functions_selefra_id | string | X | X | fk to aws_lambda_functions.selefra_id | 
| selefra_id | string | √ | √ | random id | 


