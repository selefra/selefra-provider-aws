# Table: aws_rds_cluster_parameter_group_parameters

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| account_id | string | X | √ |  | 
| description | string | X | √ |  | 
| minimum_engine_version | string | X | √ |  | 
| region | string | X | √ |  | 
| source | string | X | √ |  | 
| cluster_parameter_group_arn | string | X | √ |  | 
| allowed_values | string | X | √ |  | 
| apply_type | string | X | √ |  | 
| data_type | string | X | √ |  | 
| selefra_id | string | √ | √ | random id | 
| aws_rds_cluster_parameter_groups_selefra_id | string | X | X | fk to aws_rds_cluster_parameter_groups.selefra_id | 
| apply_method | string | X | √ |  | 
| is_modifiable | bool | X | √ |  | 
| parameter_name | string | X | √ |  | 
| parameter_value | string | X | √ |  | 
| supported_engine_modes | string_array | X | √ |  | 


