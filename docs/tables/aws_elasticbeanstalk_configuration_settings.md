# Table: aws_elasticbeanstalk_configuration_settings

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| deployment_status | string | X | √ |  | 
| application_arn | string | X | √ |  | 
| selefra_id | string | √ | √ | random id | 
| region | string | X | √ |  | 
| environment_id | string | X | √ |  | 
| date_created | timestamp | X | √ |  | 
| date_updated | timestamp | X | √ |  | 
| description | string | X | √ |  | 
| aws_elasticbeanstalk_environments_selefra_id | string | X | X | fk to aws_elasticbeanstalk_environments.selefra_id | 
| account_id | string | X | √ |  | 
| platform_arn | string | X | √ |  | 
| template_name | string | X | √ |  | 
| application_name | string | X | √ |  | 
| environment_name | string | X | √ |  | 
| option_settings | json | X | √ |  | 
| solution_stack_name | string | X | √ |  | 


