# Table: aws_docdb_cluster_snapshots

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| status | string | X | √ |  | 
| storage_encrypted | bool | X | √ |  | 
| percent_progress | int | X | √ |  | 
| snapshot_create_time | timestamp | X | √ |  | 
| aws_docdb_clusters_selefra_id | string | X | X | fk to aws_docdb_clusters.selefra_id | 
| account_id | string | X | √ |  | 
| db_cluster_identifier | string | X | √ |  | 
| engine | string | X | √ |  | 
| engine_version | string | X | √ |  | 
| port | int | X | √ |  | 
| region | string | X | √ |  | 
| cluster_create_time | timestamp | X | √ |  | 
| db_cluster_snapshot_identifier | string | X | √ |  | 
| master_username | string | X | √ |  | 
| kms_key_id | string | X | √ |  | 
| snapshot_type | string | X | √ |  | 
| source_db_cluster_snapshot_arn | string | X | √ |  | 
| vpc_id | string | X | √ |  | 
| tags | json | X | √ |  | 
| arn | string | √ | √ |  | 
| attributes | json | X | √ |  | 
| availability_zones | string_array | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 


