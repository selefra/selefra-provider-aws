# Table: aws_rds_db_snapshots

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| attributes | json | X | √ |  | 
| iops | int | X | √ |  | 
| snapshot_create_time | timestamp | X | √ |  | 
| timezone | string | X | √ |  | 
| engine_version | string | X | √ |  | 
| percent_progress | int | X | √ |  | 
| source_db_snapshot_identifier | string | X | √ |  | 
| tde_credential_arn | string | X | √ |  | 
| kms_key_id | string | X | √ |  | 
| snapshot_target | string | X | √ |  | 
| snapshot_type | string | X | √ |  | 
| encrypted | bool | X | √ |  | 
| port | int | X | √ |  | 
| engine | string | X | √ |  | 
| license_model | string | X | √ |  | 
| option_group_name | string | X | √ |  | 
| status | string | X | √ |  | 
| arn | string | √ | √ |  | 
| db_instance_identifier | string | X | √ |  | 
| db_snapshot_identifier | string | X | √ |  | 
| dbi_resource_id | string | X | √ |  | 
| storage_type | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| availability_zone | string | X | √ |  | 
| instance_create_time | timestamp | X | √ |  | 
| master_username | string | X | √ |  | 
| source_region | string | X | √ |  | 
| allocated_storage | int | X | √ |  | 
| processor_features | json | X | √ |  | 
| snapshot_database_time | timestamp | X | √ |  | 
| vpc_id | string | X | √ |  | 
| original_snapshot_create_time | timestamp | X | √ |  | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| tags | json | X | √ |  | 
| iam_database_authentication_enabled | bool | X | √ |  | 


