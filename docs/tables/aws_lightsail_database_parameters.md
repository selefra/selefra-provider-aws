# Table: aws_lightsail_database_parameters

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| apply_method | string | X | √ |  | 
| description | string | X | √ |  | 
| parameter_value | string | X | √ |  | 
| account_id | string | X | √ |  | 
| database_arn | string | X | √ |  | 
| apply_type | string | X | √ |  | 
| data_type | string | X | √ |  | 
| is_modifiable | bool | X | √ |  | 
| parameter_name | string | X | √ |  | 
| aws_lightsail_databases_selefra_id | string | X | X | fk to aws_lightsail_databases.selefra_id | 
| selefra_id | string | √ | √ | random id | 
| region | string | X | √ |  | 
| allowed_values | string | X | √ |  | 


