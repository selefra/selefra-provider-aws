# Table: aws_lambda_layer_version_policies

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| policy | string | X | √ |  | 
| revision_id | string | X | √ |  | 
| aws_lambda_layer_versions_selefra_id | string | X | X | fk to aws_lambda_layer_versions.selefra_id | 
| layer_version_arn | string | X | √ |  | 
| layer_version | int | X | √ |  | 
| result_metadata | json | X | √ |  | 
| selefra_id | string | √ | √ | random id | 


