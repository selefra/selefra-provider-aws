# Table: aws_neptune_clusters

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| engine | string | X | √ |  | 
| kms_key_id | string | X | √ |  | 
| master_username | string | X | √ |  | 
| preferred_backup_window | string | X | √ |  | 
| account_id | string | X | √ |  | 
| automatic_restart_time | timestamp | X | √ |  | 
| clone_group_id | string | X | √ |  | 
| deletion_protection | bool | X | √ |  | 
| percent_progress | string | X | √ |  | 
| backup_retention_period | int | X | √ |  | 
| enabled_cloudwatch_logs_exports | string_array | X | √ |  | 
| engine_version | string | X | √ |  | 
| multi_az | bool | X | √ |  | 
| port | int | X | √ |  | 
| vpc_security_groups | json | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| copy_tags_to_snapshot | bool | X | √ |  | 
| database_name | string | X | √ |  | 
| db_cluster_resource_id | string | X | √ |  | 
| hosted_zone_id | string | X | √ |  | 
| storage_encrypted | bool | X | √ |  | 
| preferred_maintenance_window | string | X | √ |  | 
| read_replica_identifiers | string_array | X | √ |  | 
| allocated_storage | int | X | √ |  | 
| availability_zones | string_array | X | √ |  | 
| cross_account_clone | bool | X | √ |  | 
| db_cluster_parameter_group | string | X | √ |  | 
| db_cluster_option_group_memberships | json | X | √ |  | 
| iam_database_authentication_enabled | bool | X | √ |  | 
| region | string | X | √ |  | 
| arn | string | √ | √ |  | 
| character_set_name | string | X | √ |  | 
| db_cluster_members | json | X | √ |  | 
| reader_endpoint | string | X | √ |  | 
| tags | json | X | √ |  | 
| cluster_create_time | timestamp | X | √ |  | 
| endpoint | string | X | √ |  | 
| latest_restorable_time | timestamp | X | √ |  | 
| replication_source_identifier | string | X | √ |  | 
| status | string | X | √ |  | 
| associated_roles | json | X | √ |  | 
| db_cluster_identifier | string | X | √ |  | 
| db_subnet_group | string | X | √ |  | 
| earliest_restorable_time | timestamp | X | √ |  | 


