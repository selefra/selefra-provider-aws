# Table: aws_docdb_clusters

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| db_subnet_group | string | X | √ |  | 
| engine | string | X | √ |  | 
| region | string | X | √ |  | 
| db_cluster_members | json | X | √ |  | 
| earliest_restorable_time | timestamp | X | √ |  | 
| preferred_backup_window | string | X | √ |  | 
| vpc_security_groups | json | X | √ |  | 
| arn | string | √ | √ |  | 
| associated_roles | json | X | √ |  | 
| clone_group_id | string | X | √ |  | 
| multi_az | bool | X | √ |  | 
| port | int | X | √ |  | 
| replication_source_identifier | string | X | √ |  | 
| availability_zones | string_array | X | √ |  | 
| db_cluster_resource_id | string | X | √ |  | 
| deletion_protection | bool | X | √ |  | 
| percent_progress | string | X | √ |  | 
| preferred_maintenance_window | string | X | √ |  | 
| read_replica_identifiers | string_array | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| tags | json | X | √ |  | 
| hosted_zone_id | string | X | √ |  | 
| latest_restorable_time | timestamp | X | √ |  | 
| reader_endpoint | string | X | √ |  | 
| endpoint | string | X | √ |  | 
| backup_retention_period | int | X | √ |  | 
| cluster_create_time | timestamp | X | √ |  | 
| db_cluster_identifier | string | X | √ |  | 
| db_cluster_parameter_group | string | X | √ |  | 
| enabled_cloudwatch_logs_exports | string_array | X | √ |  | 
| engine_version | string | X | √ |  | 
| kms_key_id | string | X | √ |  | 
| account_id | string | X | √ |  | 
| status | string | X | √ |  | 
| storage_encrypted | bool | X | √ |  | 
| master_username | string | X | √ |  | 


