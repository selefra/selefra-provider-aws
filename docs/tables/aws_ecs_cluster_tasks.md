# Table: aws_ecs_cluster_tasks

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| connectivity | string | X | √ |  | 
| inference_accelerators | json | X | √ |  | 
| started_by | string | X | √ |  | 
| stopped_at | timestamp | X | √ |  | 
| task_definition_arn | string | X | √ |  | 
| platform_version | string | X | √ |  | 
| pull_started_at | timestamp | X | √ |  | 
| started_at | timestamp | X | √ |  | 
| region | string | X | √ |  | 
| availability_zone | string | X | √ |  | 
| capacity_provider_name | string | X | √ |  | 
| health_status | string | X | √ |  | 
| last_status | string | X | √ |  | 
| stopping_at | timestamp | X | √ |  | 
| account_id | string | X | √ |  | 
| attachments | json | X | √ |  | 
| cpu | string | X | √ |  | 
| launch_type | string | X | √ |  | 
| containers | json | X | √ |  | 
| execution_stopped_at | timestamp | X | √ |  | 
| platform_family | string | X | √ |  | 
| stop_code | string | X | √ |  | 
| stopped_reason | string | X | √ |  | 
| container_instance_arn | string | X | √ |  | 
| desired_status | string | X | √ |  | 
| memory | string | X | √ |  | 
| aws_ecs_clusters_selefra_id | string | X | X | fk to aws_ecs_clusters.selefra_id | 
| created_at | timestamp | X | √ |  | 
| enable_execute_command | bool | X | √ |  | 
| ephemeral_storage | json | X | √ |  | 
| overrides | json | X | √ |  | 
| pull_stopped_at | timestamp | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| version | int | X | √ |  | 
| arn | string | √ | √ |  | 
| tags | json | X | √ |  | 
| attributes | json | X | √ |  | 
| cluster_arn | string | X | √ |  | 
| connectivity_at | timestamp | X | √ |  | 
| group | string | X | √ |  | 


