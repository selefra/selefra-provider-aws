# Table: aws_elasticache_clusters

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| auth_token_enabled | bool | X | √ |  | 
| preferred_outpost_arn | string | X | √ |  | 
| replication_group_log_delivery_enabled | bool | X | √ |  | 
| cache_node_type | string | X | √ |  | 
| client_download_landing_page | string | X | √ |  | 
| configuration_endpoint | json | X | √ |  | 
| auth_token_last_modified_date | timestamp | X | √ |  | 
| cache_cluster_id | string | X | √ |  | 
| preferred_availability_zone | string | X | √ |  | 
| at_rest_encryption_enabled | bool | X | √ |  | 
| cache_cluster_create_time | timestamp | X | √ |  | 
| pending_modified_values | json | X | √ |  | 
| arn | string | √ | √ |  | 
| cache_parameter_group | json | X | √ |  | 
| preferred_maintenance_window | string | X | √ |  | 
| transit_encryption_enabled | bool | X | √ |  | 
| log_delivery_configurations | json | X | √ |  | 
| notification_configuration | json | X | √ |  | 
| snapshot_retention_limit | int | X | √ |  | 
| cache_subnet_group_name | string | X | √ |  | 
| engine | string | X | √ |  | 
| num_cache_nodes | int | X | √ |  | 
| snapshot_window | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| region | string | X | √ |  | 
| cache_nodes | json | X | √ |  | 
| cache_security_groups | json | X | √ |  | 
| engine_version | string | X | √ |  | 
| replication_group_id | string | X | √ |  | 
| security_groups | json | X | √ |  | 
| account_id | string | X | √ |  | 
| auto_minor_version_upgrade | bool | X | √ |  | 
| cache_cluster_status | string | X | √ |  | 


