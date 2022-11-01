# Table: aws_ecs_cluster_container_instances

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| agent_update_status | string | X | √ |  | 
| attributes | json | X | √ |  | 
| container_instance_arn | string | X | √ |  | 
| ec2_instance_id | string | X | √ |  | 
| pending_tasks_count | int | X | √ |  | 
| remaining_resources | json | X | √ |  | 
| tags | json | X | √ |  | 
| agent_connected | bool | X | √ |  | 
| selefra_id | string | √ | √ | random id | 
| aws_ecs_clusters_selefra_id | string | X | X | fk to aws_ecs_clusters.selefra_id | 
| status | string | X | √ |  | 
| status_reason | string | X | √ |  | 
| running_tasks_count | int | X | √ |  | 
| account_id | string | X | √ |  | 
| attachments | json | X | √ |  | 
| capacity_provider_name | string | X | √ |  | 
| region | string | X | √ |  | 
| cluster_arn | string | X | √ |  | 
| registered_resources | json | X | √ |  | 
| version | int | X | √ |  | 
| version_info | json | X | √ |  | 
| health_status | json | X | √ |  | 
| registered_at | timestamp | X | √ |  | 


