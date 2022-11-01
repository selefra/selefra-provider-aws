# Table: aws_neptune_db_parameter_group_db_parameters

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| apply_type | string | X | √ |  | 
| parameter_name | string | X | √ |  | 
| parameter_value | string | X | √ |  | 
| selefra_id | string | √ | √ | random id | 
| allowed_values | string | X | √ |  | 
| description | string | X | √ |  | 
| minimum_engine_version | string | X | √ |  | 
| aws_neptune_db_parameter_groups_selefra_id | string | X | X | fk to aws_neptune_db_parameter_groups.selefra_id | 
| account_id | string | X | √ |  | 
| data_type | string | X | √ |  | 
| source | string | X | √ |  | 
| region | string | X | √ |  | 
| db_parameter_group_arn | string | X | √ |  | 
| apply_method | string | X | √ |  | 
| is_modifiable | bool | X | √ |  | 


