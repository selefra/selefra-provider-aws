# Table: aws_elbv2_target_group_target_health_descriptions

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| target_group_arn | string | X | √ |  | 
| health_check_port | string | X | √ |  | 
| target | json | X | √ |  | 
| target_health | json | X | √ |  | 
| aws_elbv2_target_groups_selefra_id | string | X | X | fk to aws_elbv2_target_groups.selefra_id | 
| selefra_id | string | √ | √ | random id | 


