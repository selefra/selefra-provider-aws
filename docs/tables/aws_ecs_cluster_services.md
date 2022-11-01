# Table: aws_ecs_cluster_services

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| cluster_arn | string | X | √ |  | 
| created_at | timestamp | X | √ |  | 
| placement_strategy | json | X | √ |  | 
| service_name | string | X | √ |  | 
| task_sets | json | X | √ |  | 
| aws_ecs_clusters_selefra_id | string | X | X | fk to aws_ecs_clusters.selefra_id | 
| account_id | string | X | √ |  | 
| deployment_configuration | json | X | √ |  | 
| enable_execute_command | bool | X | √ |  | 
| load_balancers | json | X | √ |  | 
| network_configuration | json | X | √ |  | 
| placement_constraints | json | X | √ |  | 
| platform_family | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| created_by | string | X | √ |  | 
| deployments | json | X | √ |  | 
| launch_type | string | X | √ |  | 
| pending_count | int | X | √ |  | 
| role_arn | string | X | √ |  | 
| service_registries | json | X | √ |  | 
| task_definition | string | X | √ |  | 
| tags | json | X | √ |  | 
| platform_version | string | X | √ |  | 
| region | string | X | √ |  | 
| arn | string | √ | √ |  | 
| desired_count | int | X | √ |  | 
| events | json | X | √ |  | 
| propagate_tags | string | X | √ |  | 
| scheduling_strategy | string | X | √ |  | 
| deployment_controller | json | X | √ |  | 
| enable_ecs_managed_tags | bool | X | √ |  | 
| health_check_grace_period_seconds | int | X | √ |  | 
| running_count | int | X | √ |  | 
| capacity_provider_strategy | json | X | √ |  | 
| status | string | X | √ |  | 


