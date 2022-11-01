# Table: aws_elasticache_replication_groups

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| automatic_failover | string | X | √ |  | 
| log_delivery_configurations | json | X | √ |  | 
| member_clusters | string_array | X | √ |  | 
| snapshot_window | string | X | √ |  | 
| status | string | X | √ |  | 
| transit_encryption_enabled | bool | X | √ |  | 
| at_rest_encryption_enabled | bool | X | √ |  | 
| auth_token_enabled | bool | X | √ |  | 
| node_groups | json | X | √ |  | 
| region | string | X | √ |  | 
| cache_node_type | string | X | √ |  | 
| configuration_endpoint | json | X | √ |  | 
| multi_az | string | X | √ |  | 
| member_clusters_outpost_arns | string_array | X | √ |  | 
| snapshot_retention_limit | int | X | √ |  | 
| user_group_ids | string_array | X | √ |  | 
| auto_minor_version_upgrade | bool | X | √ |  | 
| kms_key_id | string | X | √ |  | 
| global_replication_group_info | json | X | √ |  | 
| account_id | string | X | √ |  | 
| cluster_enabled | bool | X | √ |  | 
| replication_group_create_time | timestamp | X | √ |  | 
| snapshotting_cluster_id | string | X | √ |  | 
| pending_modified_values | json | X | √ |  | 
| data_tiering | string | X | √ |  | 
| description | string | X | √ |  | 
| replication_group_id | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| arn | string | √ | √ |  | 
| auth_token_last_modified_date | timestamp | X | √ |  | 


