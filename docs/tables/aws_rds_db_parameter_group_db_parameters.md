# Table: aws_rds_db_parameter_group_db_parameters

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| parameter_value | string | X | √ |  | 
| account_id | string | X | √ |  | 
| allowed_values | string | X | √ |  | 
| apply_method | string | X | √ |  | 
| selefra_id | string | √ | √ | random id | 
| aws_rds_db_parameter_groups_selefra_id | string | X | X | fk to aws_rds_db_parameter_groups.selefra_id | 
| description | string | X | √ |  | 
| minimum_engine_version | string | X | √ |  | 
| supported_engine_modes | string_array | X | √ |  | 
| data_type | string | X | √ |  | 
| is_modifiable | bool | X | √ |  | 
| parameter_name | string | X | √ |  | 
| source | string | X | √ |  | 
| region | string | X | √ |  | 
| db_parameter_group_arn | string | X | √ |  | 
| apply_type | string | X | √ |  | 


