# Table: aws_rds_cluster_snapshots

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| engine | string | X | √ |  | 
| engine_mode | string | X | √ |  | 
| kms_key_id | string | X | √ |  | 
| master_username | string | X | √ |  | 
| percent_progress | int | X | √ |  | 
| tags | json | X | √ |  | 
| allocated_storage | int | X | √ |  | 
| availability_zones | string_array | X | √ |  | 
| snapshot_type | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| arn | string | √ | √ |  | 
| cluster_create_time | timestamp | X | √ |  | 
| vpc_id | string | X | √ |  | 
| storage_encrypted | bool | X | √ |  | 
| region | string | X | √ |  | 
| port | int | X | √ |  | 
| status | string | X | √ |  | 
| db_cluster_snapshot_identifier | string | X | √ |  | 
| engine_version | string | X | √ |  | 
| iam_database_authentication_enabled | bool | X | √ |  | 
| license_model | string | X | √ |  | 
| snapshot_create_time | timestamp | X | √ |  | 
| account_id | string | X | √ |  | 
| attributes | json | X | √ |  | 
| db_cluster_identifier | string | X | √ |  | 
| source_db_cluster_snapshot_arn | string | X | √ |  | 


