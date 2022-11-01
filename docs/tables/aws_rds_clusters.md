# Table: aws_rds_clusters

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| db_cluster_members | json | X | √ |  | 
| performance_insights_enabled | bool | X | √ |  | 
| tags | json | X | √ |  | 
| allocated_storage | int | X | √ |  | 
| capacity | int | X | √ |  | 
| copy_tags_to_snapshot | bool | X | √ |  | 
| domain_memberships | json | X | √ |  | 
| hosted_zone_id | string | X | √ |  | 
| performance_insights_retention_period | int | X | √ |  | 
| status | string | X | √ |  | 
| activity_stream_kinesis_stream_name | string | X | √ |  | 
| associated_roles | json | X | √ |  | 
| arn | string | √ | √ |  | 
| activity_stream_status | string | X | √ |  | 
| performance_insights_kms_key_id | string | X | √ |  | 
| preferred_maintenance_window | string | X | √ |  | 
| publicly_accessible | bool | X | √ |  | 
| storage_encrypted | bool | X | √ |  | 
| storage_type | string | X | √ |  | 
| availability_zones | string_array | X | √ |  | 
| cross_account_clone | bool | X | √ |  | 
| db_cluster_parameter_group | string | X | √ |  | 
| preferred_backup_window | string | X | √ |  | 
| vpc_security_groups | json | X | √ |  | 
| account_id | string | X | √ |  | 
| endpoint | string | X | √ |  | 
| http_endpoint_enabled | bool | X | √ |  | 
| scaling_configuration_info | json | X | √ |  | 
| engine_mode | string | X | √ |  | 
| multi_az | bool | X | √ |  | 
| iam_database_authentication_enabled | bool | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| percent_progress | string | X | √ |  | 
| clone_group_id | string | X | √ |  | 
| deletion_protection | bool | X | √ |  | 
| engine_version | string | X | √ |  | 
| monitoring_role_arn | string | X | √ |  | 
| serverless_v2_scaling_configuration | json | X | √ |  | 
| backtrack_window | int | X | √ |  | 
| db_cluster_identifier | string | X | √ |  | 
| database_name | string | X | √ |  | 
| db_cluster_resource_id | string | X | √ |  | 
| network_type | string | X | √ |  | 
| port | int | X | √ |  | 
| reader_endpoint | string | X | √ |  | 
| cluster_create_time | timestamp | X | √ |  | 
| earliest_backtrack_time | timestamp | X | √ |  | 
| enabled_cloudwatch_logs_exports | string_array | X | √ |  | 
| master_username | string | X | √ |  | 
| kms_key_id | string | X | √ |  | 
| replication_source_identifier | string | X | √ |  | 
| activity_stream_kms_key_id | string | X | √ |  | 
| automatic_restart_time | timestamp | X | √ |  | 
| character_set_name | string | X | √ |  | 
| engine | string | X | √ |  | 
| db_subnet_group | string | X | √ |  | 
| earliest_restorable_time | timestamp | X | √ |  | 
| global_write_forwarding_requested | bool | X | √ |  | 
| global_write_forwarding_status | string | X | √ |  | 
| activity_stream_mode | string | X | √ |  | 
| backup_retention_period | int | X | √ |  | 
| custom_endpoints | string_array | X | √ |  | 
| db_cluster_option_group_memberships | json | X | √ |  | 
| iops | int | X | √ |  | 
| latest_restorable_time | timestamp | X | √ |  | 
| read_replica_identifiers | string_array | X | √ |  | 
| db_cluster_instance_class | string | X | √ |  | 
| pending_modified_values | json | X | √ |  | 
| region | string | X | √ |  | 
| auto_minor_version_upgrade | bool | X | √ |  | 
| backtrack_consumed_change_records | int | X | √ |  | 
| monitoring_interval | int | X | √ |  | 


