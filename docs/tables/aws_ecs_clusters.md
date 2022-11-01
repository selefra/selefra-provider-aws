# Table: aws_ecs_clusters

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| attachments | json | X | √ |  | 
| attachments_status | string | X | √ |  | 
| configuration | json | X | √ |  | 
| default_capacity_provider_strategy | json | X | √ |  | 
| arn | string | √ | √ |  | 
| tags | json | X | √ |  | 
| active_services_count | int | X | √ |  | 
| account_id | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| status | string | X | √ |  | 
| capacity_providers | string_array | X | √ |  | 
| cluster_name | string | X | √ |  | 
| pending_tasks_count | int | X | √ |  | 
| registered_container_instances_count | int | X | √ |  | 
| running_tasks_count | int | X | √ |  | 
| settings | json | X | √ |  | 
| statistics | json | X | √ |  | 
| region | string | X | √ |  | 


