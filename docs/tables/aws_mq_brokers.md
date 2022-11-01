# Table: aws_mq_brokers

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| region | string | X | √ |  | 
| authentication_strategy | string | X | √ |  | 
| ldap_server_metadata | json | X | √ |  | 
| logs | json | X | √ |  | 
| publicly_accessible | bool | X | √ |  | 
| security_groups | string_array | X | √ |  | 
| storage_type | string | X | √ |  | 
| users | json | X | √ |  | 
| encryption_options | json | X | √ |  | 
| engine_version | string | X | √ |  | 
| pending_authentication_strategy | string | X | √ |  | 
| pending_ldap_server_metadata | json | X | √ |  | 
| arn | string | √ | √ |  | 
| configurations | json | X | √ |  | 
| result_metadata | json | X | √ |  | 
| broker_name | string | X | √ |  | 
| broker_state | string | X | √ |  | 
| host_instance_type | string | X | √ |  | 
| pending_host_instance_type | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| actions_required | json | X | √ |  | 
| deployment_mode | string | X | √ |  | 
| engine_type | string | X | √ |  | 
| subnet_ids | string_array | X | √ |  | 
| auto_minor_version_upgrade | bool | X | √ |  | 
| broker_id | string | X | √ |  | 
| maintenance_window_start_time | json | X | √ |  | 
| tags | json | X | √ |  | 
| account_id | string | X | √ |  | 
| broker_instances | json | X | √ |  | 
| created | timestamp | X | √ |  | 
| pending_engine_version | string | X | √ |  | 
| pending_security_groups | string_array | X | √ |  | 


