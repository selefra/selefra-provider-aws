# Table: aws_autoscaling_group_lifecycle_hooks

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| lifecycle_transition | string | X | √ |  | 
| notification_target_arn | string | X | √ |  | 
| role_arn | string | X | √ |  | 
| selefra_id | string | √ | √ | random id | 
| account_id | string | X | √ |  | 
| heartbeat_timeout | int | X | √ |  | 
| lifecycle_hook_name | string | X | √ |  | 
| region | string | X | √ |  | 
| auto_scaling_group_name | string | X | √ |  | 
| global_timeout | int | X | √ |  | 
| aws_autoscaling_groups_selefra_id | string | X | X | fk to aws_autoscaling_groups.selefra_id | 
| group_arn | string | X | √ |  | 
| default_result | string | X | √ |  | 
| notification_metadata | string | X | √ |  | 


