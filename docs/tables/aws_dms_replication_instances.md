# Table: aws_dms_replication_instances

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| dns_name_servers | string | X | √ |  | 
| replication_instance_class | string | X | √ |  | 
| replication_instance_public_ip_addresses | string_array | X | √ |  | 
| secondary_availability_zone | string | X | √ |  | 
| instance_create_time | timestamp | X | √ |  | 
| replication_instance_public_ip_address | string | X | √ |  | 
| tags | json | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| region | string | X | √ |  | 
| arn | string | √ | √ |  | 
| auto_minor_version_upgrade | bool | X | √ |  | 
| engine_version | string | X | √ |  | 
| publicly_accessible | bool | X | √ |  | 
| replication_instance_identifier | string | X | √ |  | 
| vpc_security_groups | json | X | √ |  | 
| account_id | string | X | √ |  | 
| multi_az | bool | X | √ |  | 
| pending_modified_values | json | X | √ |  | 
| replication_instance_private_ip_addresses | string_array | X | √ |  | 
| replication_instance_status | string | X | √ |  | 
| replication_subnet_group | json | X | √ |  | 
| allocated_storage | int | X | √ |  | 
| availability_zone | string | X | √ |  | 
| free_until | timestamp | X | √ |  | 
| kms_key_id | string | X | √ |  | 
| preferred_maintenance_window | string | X | √ |  | 
| replication_instance_private_ip_address | string | X | √ |  | 


