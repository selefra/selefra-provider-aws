# Table: aws_glue_registry_schema_versions

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| result_metadata | json | X | √ |  | 
| aws_glue_registry_schemas_selefra_id | string | X | X | fk to aws_glue_registry_schemas.selefra_id | 
| version_number | int | X | √ |  | 
| account_id | string | X | √ |  | 
| metadata | json | X | √ |  | 
| schema_definition | string | X | √ |  | 
| status | string | X | √ |  | 
| selefra_id | string | √ | √ | random id | 
| region | string | X | √ |  | 
| registry_schema_arn | string | X | √ |  | 
| created_time | string | X | √ |  | 
| data_format | string | X | √ |  | 
| schema_arn | string | X | √ |  | 
| schema_version_id | string | X | √ |  | 


