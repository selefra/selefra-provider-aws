# Table: aws_lambda_function_concurrency_configs

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| account_id | string | X | √ |  | 
| allocated_provisioned_concurrent_executions | int | X | √ |  | 
| last_modified | string | X | √ |  | 
| status | string | X | √ |  | 
| aws_lambda_functions_selefra_id | string | X | X | fk to aws_lambda_functions.selefra_id | 
| selefra_id | string | √ | √ | random id | 
| region | string | X | √ |  | 
| function_arn | string | X | √ |  | 
| available_provisioned_concurrent_executions | int | X | √ |  | 
| requested_provisioned_concurrent_executions | int | X | √ |  | 
| status_reason | string | X | √ |  | 


