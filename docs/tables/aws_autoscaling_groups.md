# Table: aws_autoscaling_groups

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| account_id | string | X | √ |  | 
| tags | json | X | √ |  | 
| min_size | int | X | √ |  | 
| health_check_grace_period | int | X | √ |  | 
| service_linked_role_arn | string | X | √ |  | 
| status | string | X | √ |  | 
| target_group_ar_ns | string_array | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| created_time | timestamp | X | √ |  | 
| health_check_type | string | X | √ |  | 
| default_instance_warmup | int | X | √ |  | 
| enabled_metrics | json | X | √ |  | 
| instances | json | X | √ |  | 
| new_instances_protected_from_scale_in | bool | X | √ |  | 
| max_size | int | X | √ |  | 
| launch_template | json | X | √ |  | 
| mixed_instances_policy | json | X | √ |  | 
| placement_group | string | X | √ |  | 
| predicted_capacity | int | X | √ |  | 
| availability_zones | string_array | X | √ |  | 
| default_cooldown | int | X | √ |  | 
| warm_pool_size | int | X | √ |  | 
| load_balancers | json | X | √ |  | 
| context | string | X | √ |  | 
| launch_configuration_name | string | X | √ |  | 
| load_balancer_names | string_array | X | √ |  | 
| suspended_processes | json | X | √ |  | 
| arn | string | √ | √ |  | 
| auto_scaling_group_name | string | X | √ |  | 
| capacity_rebalance | bool | X | √ |  | 
| desired_capacity_type | string | X | √ |  | 
| termination_policies | string_array | X | √ |  | 
| vpc_zone_identifier | string | X | √ |  | 
| region | string | X | √ |  | 
| load_balancer_target_groups | json | X | √ |  | 
| desired_capacity | int | X | √ |  | 
| max_instance_lifetime | int | X | √ |  | 
| warm_pool_configuration | json | X | √ |  | 
| notification_configurations | json | X | √ |  | 


