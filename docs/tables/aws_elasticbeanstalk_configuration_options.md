# Table: aws_elasticbeanstalk_configuration_options

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| max_value | int | X | √ |  | 
| namespace | string | X | √ |  | 
| user_defined | bool | X | √ |  | 
| account_id | string | X | √ |  | 
| environment_id | string | X | √ |  | 
| value_options | string_array | X | √ |  | 
| aws_elasticbeanstalk_environments_selefra_id | string | X | X | fk to aws_elasticbeanstalk_environments.selefra_id | 
| region | string | X | √ |  | 
| max_length | int | X | √ |  | 
| name | string | X | √ |  | 
| application_arn | string | X | √ |  | 
| selefra_id | string | √ | √ | random id | 
| change_severity | string | X | √ |  | 
| default_value | string | X | √ |  | 
| min_value | int | X | √ |  | 
| regex | json | X | √ |  | 
| value_type | string | X | √ |  | 


