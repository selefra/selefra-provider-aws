# Table: aws_ecs_task_definitions

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| status | string | X | √ |  | 
| cpu | string | X | √ |  | 
| family | string | X | √ |  | 
| network_mode | string | X | √ |  | 
| requires_attributes | json | X | √ |  | 
| runtime_platform | json | X | √ |  | 
| region | string | X | √ |  | 
| compatibilities | string_array | X | √ |  | 
| ipc_mode | string | X | √ |  | 
| tags | json | X | √ |  | 
| inference_accelerators | json | X | √ |  | 
| requires_compatibilities | string_array | X | √ |  | 
| revision | int | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| container_definitions | json | X | √ |  | 
| ephemeral_storage | json | X | √ |  | 
| pid_mode | string | X | √ |  | 
| task_role_arn | string | X | √ |  | 
| account_id | string | X | √ |  | 
| arn | string | √ | √ |  | 
| placement_constraints | json | X | √ |  | 
| registered_at | timestamp | X | √ |  | 
| volumes | json | X | √ |  | 
| deregistered_at | timestamp | X | √ |  | 
| execution_role_arn | string | X | √ |  | 
| memory | string | X | √ |  | 
| registered_by | string | X | √ |  | 
| proxy_configuration | json | X | √ |  | 


