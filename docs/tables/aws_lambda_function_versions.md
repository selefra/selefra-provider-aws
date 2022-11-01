# Table: aws_lambda_function_versions

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| dead_letter_config | json | X | √ |  | 
| layers | json | X | √ |  | 
| revision_id | string | X | √ |  | 
| tracing_config | json | X | √ |  | 
| vpc_config | json | X | √ |  | 
| selefra_id | string | √ | √ | random id | 
| file_system_configs | json | X | √ |  | 
| handler | string | X | √ |  | 
| signing_job_arn | string | X | √ |  | 
| memory_size | int | X | √ |  | 
| role | string | X | √ |  | 
| signing_profile_version_arn | string | X | √ |  | 
| state_reason | string | X | √ |  | 
| timeout | int | X | √ |  | 
| account_id | string | X | √ |  | 
| function_arn | string | X | √ |  | 
| architectures | string_array | X | √ |  | 
| master_arn | string | X | √ |  | 
| state | string | X | √ |  | 
| aws_lambda_functions_selefra_id | string | X | X | fk to aws_lambda_functions.selefra_id | 
| code_sha256 | string | X | √ |  | 
| ephemeral_storage | json | X | √ |  | 
| description | string | X | √ |  | 
| environment | json | X | √ |  | 
| last_modified | string | X | √ |  | 
| runtime | string | X | √ |  | 
| state_reason_code | string | X | √ |  | 
| version | string | X | √ |  | 
| region | string | X | √ |  | 
| image_config_response | json | X | √ |  | 
| package_type | string | X | √ |  | 
| code_size | int | X | √ |  | 
| function_name | string | X | √ |  | 
| kms_key_arn | string | X | √ |  | 
| last_update_status | string | X | √ |  | 
| last_update_status_reason | string | X | √ |  | 
| last_update_status_reason_code | string | X | √ |  | 


