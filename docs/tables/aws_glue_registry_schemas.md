# Table: aws_glue_registry_schemas

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| latest_schema_version | int | X | √ |  | 
| next_schema_version | int | X | √ |  | 
| registry_name | string | X | √ |  | 
| schema_status | string | X | √ |  | 
| updated_time | string | X | √ |  | 
| created_time | string | X | √ |  | 
| description | string | X | √ |  | 
| schema_name | string | X | √ |  | 
| aws_glue_registries_selefra_id | string | X | X | fk to aws_glue_registries.selefra_id | 
| arn | string | X | √ |  | 
| tags | json | X | √ |  | 
| schema_checkpoint | int | X | √ |  | 
| selefra_id | string | √ | √ | random id | 
| data_format | string | X | √ |  | 
| registry_arn | string | X | √ |  | 
| compatibility | string | X | √ |  | 
| result_metadata | json | X | √ |  | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 


