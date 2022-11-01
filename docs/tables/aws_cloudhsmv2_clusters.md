# Table: aws_cloudhsmv2_clusters

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| backup_retention_policy | json | X | √ |  | 
| cluster_id | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| pre_co_password | string | X | √ |  | 
| arn | string | √ | √ |  | 
| tags | json | X | √ |  | 
| security_group | string | X | √ |  | 
| source_backup_id | string | X | √ |  | 
| subnet_mapping | json | X | √ |  | 
| hsms | json | X | √ |  | 
| state | string | X | √ |  | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| backup_policy | string | X | √ |  | 
| certificates | json | X | √ |  | 
| create_timestamp | timestamp | X | √ |  | 
| hsm_type | string | X | √ |  | 
| state_message | string | X | √ |  | 
| vpc_id | string | X | √ |  | 


