# Table: aws_neptune_instances

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| status_infos | json | X | √ |  | 
| timezone | string | X | √ |  | 
| engine | string | X | √ |  | 
| latest_restorable_time | timestamp | X | √ |  | 
| preferred_maintenance_window | string | X | √ |  | 
| kms_key_id | string | X | √ |  | 
| multi_az | bool | X | √ |  | 
| publicly_accessible | bool | X | √ |  | 
| iops | int | X | √ |  | 
| enhanced_monitoring_resource_arn | string | X | √ |  | 
| secondary_availability_zone | string | X | √ |  | 
| db_instance_class | string | X | √ |  | 
| db_subnet_group | json | X | √ |  | 
| monitoring_interval | int | X | √ |  | 
| monitoring_role_arn | string | X | √ |  | 
| account_id | string | X | √ |  | 
| enabled_cloudwatch_logs_exports | string_array | X | √ |  | 
| master_username | string | X | √ |  | 
| endpoint | json | X | √ |  | 
| engine_version | string | X | √ |  | 
| preferred_backup_window | string | X | √ |  | 
| tde_credential_arn | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| backup_retention_period | int | X | √ |  | 
| db_cluster_identifier | string | X | √ |  | 
| deletion_protection | bool | X | √ |  | 
| option_group_memberships | json | X | √ |  | 
| read_replica_db_cluster_identifiers | string_array | X | √ |  | 
| ca_certificate_identifier | string | X | √ |  | 
| storage_type | string | X | √ |  | 
| license_model | string | X | √ |  | 
| performance_insights_enabled | bool | X | √ |  | 
| region | string | X | √ |  | 
| db_instance_identifier | string | X | √ |  | 
| db_instance_port | int | X | √ |  | 
| dbi_resource_id | string | X | √ |  | 
| instance_create_time | timestamp | X | √ |  | 
| tags | json | X | √ |  | 
| auto_minor_version_upgrade | bool | X | √ |  | 
| db_instance_status | string | X | √ |  | 
| db_name | string | X | √ |  | 
| db_security_groups | json | X | √ |  | 
| performance_insights_kms_key_id | string | X | √ |  | 
| read_replica_db_instance_identifiers | string_array | X | √ |  | 
| read_replica_source_db_instance_identifier | string | X | √ |  | 
| vpc_security_groups | json | X | √ |  | 
| arn | string | √ | √ |  | 
| allocated_storage | int | X | √ |  | 
| iam_database_authentication_enabled | bool | X | √ |  | 
| pending_modified_values | json | X | √ |  | 
| availability_zone | string | X | √ |  | 
| character_set_name | string | X | √ |  | 
| domain_memberships | json | X | √ |  | 
| storage_encrypted | bool | X | √ |  | 
| copy_tags_to_snapshot | bool | X | √ |  | 
| db_parameter_groups | json | X | √ |  | 
| promotion_tier | int | X | √ |  | 


