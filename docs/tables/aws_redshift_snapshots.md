# Table: aws_redshift_snapshots

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| current_backup_rate_in_mega_bytes_per_second | float | X | √ |  | 
| engine_full_version | string | X | √ |  | 
| manual_snapshot_remaining_days | int | X | √ |  | 
| owner_account | string | X | √ |  | 
| encrypted | bool | X | √ |  | 
| encrypted_with_hsm | bool | X | √ |  | 
| kms_key_id | string | X | √ |  | 
| source_region | string | X | √ |  | 
| tags | json | X | √ | `Tags consisting of a name/value pair for a resource.` | 
| db_name | string | X | √ |  | 
| total_backup_size_in_mega_bytes | float | X | √ |  | 
| region | string | X | √ |  | 
| master_username | string | X | √ |  | 
| aws_redshift_clusters_selefra_id | string | X | X | fk to aws_redshift_clusters.selefra_id | 
| manual_snapshot_retention_period | int | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| actual_incremental_backup_size_in_mega_bytes | float | X | √ |  | 
| cluster_version | string | X | √ |  | 
| enhanced_vpc_routing | bool | X | √ |  | 
| maintenance_track_name | string | X | √ |  | 
| snapshot_identifier | string | X | √ |  | 
| snapshot_retention_start_time | timestamp | X | √ |  | 
| snapshot_type | string | X | √ |  | 
| backup_progress_in_mega_bytes | float | X | √ |  | 
| cluster_create_time | timestamp | X | √ |  | 
| node_type | string | X | √ |  | 
| port | int | X | √ |  | 
| number_of_nodes | int | X | √ |  | 
| restorable_node_types | string_array | X | √ |  | 
| snapshot_create_time | timestamp | X | √ |  | 
| status | string | X | √ |  | 
| account_id | string | X | √ |  | 
| arn | string | √ | √ | `ARN of the snapshot.` | 
| availability_zone | string | X | √ |  | 
| cluster_identifier | string | X | √ |  | 
| vpc_id | string | X | √ |  | 
| accounts_with_restore_access | json | X | √ |  | 
| elapsed_time_in_seconds | int | X | √ |  | 
| estimated_seconds_to_completion | int | X | √ |  | 


