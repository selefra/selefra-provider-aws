# Table: aws_cloudhsmv2_backups

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| copy_timestamp | timestamp | X | √ |  | 
| delete_timestamp | timestamp | X | √ |  | 
| source_region | string | X | √ |  | 
| source_backup | string | X | √ |  | 
| region | string | X | √ |  | 
| create_timestamp | timestamp | X | √ |  | 
| never_expires | bool | X | √ |  | 
| tags | json | X | √ |  | 
| cluster_id | string | X | √ |  | 
| backup_state | string | X | √ |  | 
| source_cluster | string | X | √ |  | 
| account_id | string | X | √ |  | 
| arn | string | √ | √ |  | 
| backup_id | string | X | √ |  | 


