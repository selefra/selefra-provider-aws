# Table: aws_neptune_cluster_snapshots

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| vpc_id | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| arn | string | √ | √ |  | 
| attributes | json | X | √ |  | 
| allocated_storage | int | X | √ |  | 
| db_cluster_identifier | string | X | √ |  | 
| license_model | string | X | √ |  | 
| status | string | X | √ |  | 
| availability_zones | string_array | X | √ |  | 
| cluster_create_time | timestamp | X | √ |  | 
| kms_key_id | string | X | √ |  | 
| snapshot_create_time | timestamp | X | √ |  | 
| snapshot_type | string | X | √ |  | 
| tags | json | X | √ |  | 
| engine_version | string | X | √ |  | 
| iam_database_authentication_enabled | bool | X | √ |  | 
| percent_progress | int | X | √ |  | 
| storage_encrypted | bool | X | √ |  | 
| source_db_cluster_snapshot_arn | string | X | √ |  | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| db_cluster_snapshot_identifier | string | X | √ |  | 
| engine | string | X | √ |  | 
| master_username | string | X | √ |  | 
| port | int | X | √ |  | 


