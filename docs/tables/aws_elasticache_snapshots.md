# Table: aws_elasticache_snapshots

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| arn | string | √ | √ |  | 
| preferred_availability_zone | string | X | √ |  | 
| replication_group_description | string | X | √ |  | 
| topic_arn | string | X | √ |  | 
| auto_minor_version_upgrade | bool | X | √ |  | 
| automatic_failover | string | X | √ |  | 
| cache_subnet_group_name | string | X | √ |  | 
| kms_key_id | string | X | √ |  | 
| snapshot_source | string | X | √ |  | 
| snapshot_status | string | X | √ |  | 
| engine | string | X | √ |  | 
| num_cache_nodes | int | X | √ |  | 
| num_node_groups | int | X | √ |  | 
| cache_parameter_group_name | string | X | √ |  | 
| port | int | X | √ |  | 
| replication_group_id | string | X | √ |  | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| cache_node_type | string | X | √ |  | 
| preferred_maintenance_window | string | X | √ |  | 
| preferred_outpost_arn | string | X | √ |  | 
| snapshot_retention_limit | int | X | √ |  | 
| snapshot_window | string | X | √ |  | 
| vpc_id | string | X | √ |  | 
| cache_cluster_create_time | timestamp | X | √ |  | 
| cache_cluster_id | string | X | √ |  | 
| data_tiering | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| snapshot_name | string | X | √ |  | 
| engine_version | string | X | √ |  | 
| node_snapshots | json | X | √ |  | 


