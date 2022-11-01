# Table: aws_rds_instances

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| dbi_resource_id | string | X | √ |  | 
| engine | string | X | √ |  | 
| activity_stream_kinesis_stream_name | string | X | √ |  | 
| db_instance_status | string | X | √ |  | 
| db_name | string | X | √ |  | 
| pending_modified_values | json | X | √ |  | 
| db_instance_identifier | string | X | √ |  | 
| instance_create_time | timestamp | X | √ |  | 
| kms_key_id | string | X | √ |  | 
| db_security_groups | json | X | √ |  | 
| iam_database_authentication_enabled | bool | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| activity_stream_mode | string | X | √ |  | 
| associated_roles | json | X | √ |  | 
| availability_zone | string | X | √ |  | 
| option_group_memberships | json | X | √ |  | 
| preferred_backup_window | string | X | √ |  | 
| resume_full_automation_mode_time | timestamp | X | √ |  | 
| automatic_restart_time | timestamp | X | √ |  | 
| backup_retention_period | int | X | √ |  | 
| latest_restorable_time | timestamp | X | √ |  | 
| deletion_protection | bool | X | √ |  | 
| domain_memberships | json | X | √ |  | 
| activity_stream_kms_key_id | string | X | √ |  | 
| automation_mode | string | X | √ |  | 
| db_parameter_groups | json | X | √ |  | 
| aws_backup_recovery_point_arn | string | X | √ |  | 
| backup_target | string | X | √ |  | 
| timezone | string | X | √ |  | 
| arn | string | √ | √ |  | 
| processor_features | json | X | √ |  | 
| vpc_security_groups | json | X | √ |  | 
| engine_version | string | X | √ |  | 
| max_allocated_storage | int | X | √ |  | 
| activity_stream_policy_status | string | X | √ |  | 
| auto_minor_version_upgrade | bool | X | √ |  | 
| db_subnet_group | json | X | √ |  | 
| license_model | string | X | √ |  | 
| listener_endpoint | json | X | √ |  | 
| multi_az | bool | X | √ |  | 
| performance_insights_enabled | bool | X | √ |  | 
| read_replica_db_instance_identifiers | string_array | X | √ |  | 
| activity_stream_engine_native_audit_fields_included | bool | X | √ |  | 
| activity_stream_status | string | X | √ |  | 
| copy_tags_to_snapshot | bool | X | √ |  | 
| monitoring_role_arn | string | X | √ |  | 
| performance_insights_kms_key_id | string | X | √ |  | 
| preferred_maintenance_window | string | X | √ |  | 
| storage_type | string | X | √ |  | 
| tde_credential_arn | string | X | √ |  | 
| region | string | X | √ |  | 
| ca_certificate_identifier | string | X | √ |  | 
| db_instance_class | string | X | √ |  | 
| account_id | string | X | √ |  | 
| character_set_name | string | X | √ |  | 
| endpoint | json | X | √ |  | 
| customer_owned_ip_enabled | bool | X | √ |  | 
| performance_insights_retention_period | int | X | √ |  | 
| storage_encrypted | bool | X | √ |  | 
| nchar_character_set_name | string | X | √ |  | 
| publicly_accessible | bool | X | √ |  | 
| secondary_availability_zone | string | X | √ |  | 
| tags | json | X | √ |  | 
| db_cluster_identifier | string | X | √ |  | 
| enhanced_monitoring_resource_arn | string | X | √ |  | 
| iops | int | X | √ |  | 
| promotion_tier | int | X | √ |  | 
| custom_iam_instance_profile | string | X | √ |  | 
| db_instance_automated_backups_replications | json | X | √ |  | 
| db_instance_port | int | X | √ |  | 
| network_type | string | X | √ |  | 
| replica_mode | string | X | √ |  | 
| enabled_cloudwatch_logs_exports | string_array | X | √ |  | 
| master_username | string | X | √ |  | 
| monitoring_interval | int | X | √ |  | 
| status_infos | json | X | √ |  | 
| allocated_storage | int | X | √ |  | 
| read_replica_db_cluster_identifiers | string_array | X | √ |  | 
| read_replica_source_db_instance_identifier | string | X | √ |  | 


