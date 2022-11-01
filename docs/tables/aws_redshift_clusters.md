# Table: aws_redshift_clusters

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| arn | string | √ | √ | `The Amazon Resource Name (ARN) for the resource.` | 
| automated_snapshot_retention_period | int | X | √ |  | 
| availability_zone_relocation_status | string | X | √ |  | 
| number_of_nodes | int | X | √ |  | 
| snapshot_schedule_state | string | X | √ |  | 
| logging_status | json | X | √ | `Describes the status of logging for a cluster.` | 
| default_iam_role_arn | string | X | √ |  | 
| cluster_public_key | string | X | √ |  | 
| iam_roles | json | X | √ |  | 
| cluster_create_time | timestamp | X | √ |  | 
| data_transfer_progress | json | X | √ |  | 
| modify_status | string | X | √ |  | 
| preferred_maintenance_window | string | X | √ |  | 
| reserved_node_exchange_status | json | X | √ |  | 
| db_name | string | X | √ |  | 
| node_type | string | X | √ |  | 
| vpc_id | string | X | √ |  | 
| account_id | string | X | √ |  | 
| elastic_ip_status | json | X | √ |  | 
| expected_next_snapshot_schedule_time | timestamp | X | √ |  | 
| publicly_accessible | bool | X | √ |  | 
| cluster_snapshot_copy_status | json | X | √ |  | 
| elastic_resize_number_of_node_options | string | X | √ |  | 
| manual_snapshot_retention_period | int | X | √ |  | 
| snapshot_schedule_identifier | string | X | √ |  | 
| cluster_status | string | X | √ |  | 
| deferred_maintenance_windows | json | X | √ |  | 
| expected_next_snapshot_schedule_time_status | string | X | √ |  | 
| restore_status | json | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| availability_zone | string | X | √ |  | 
| encrypted | bool | X | √ |  | 
| endpoint | json | X | √ |  | 
| enhanced_vpc_routing | bool | X | √ |  | 
| pending_actions | string_array | X | √ |  | 
| vpc_security_groups | json | X | √ |  | 
| cluster_namespace_arn | string | X | √ |  | 
| cluster_revision_number | string | X | √ |  | 
| cluster_subnet_group_name | string | X | √ |  | 
| allow_version_upgrade | bool | X | √ |  | 
| maintenance_track_name | string | X | √ |  | 
| master_username | string | X | √ |  | 
| total_storage_capacity_in_mega_bytes | int | X | √ |  | 
| cluster_availability_status | string | X | √ |  | 
| hsm_status | json | X | √ |  | 
| aqua_configuration | json | X | √ |  | 
| pending_modified_values | json | X | √ |  | 
| tags | json | X | √ | `The list of tags for the cluster.` | 
| cluster_version | string | X | √ |  | 
| next_maintenance_window_start_time | timestamp | X | √ |  | 
| region | string | X | √ |  | 
| cluster_identifier | string | X | √ |  | 
| cluster_nodes | json | X | √ |  | 
| resize_info | json | X | √ |  | 
| cluster_security_groups | json | X | √ |  | 
| kms_key_id | string | X | √ |  | 


