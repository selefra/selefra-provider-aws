# Table: aws_neptune_global_clusters

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| global_cluster_identifier | string | X | √ |  | 
| global_cluster_members | json | X | √ |  | 
| global_cluster_resource_id | string | X | √ |  | 
| status | string | X | √ |  | 
| region | string | X | √ |  | 
| arn | string | √ | √ |  | 
| engine | string | X | √ |  | 
| engine_version | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| account_id | string | X | √ |  | 
| tags | json | X | √ |  | 
| deletion_protection | bool | X | √ |  | 
| storage_encrypted | bool | X | √ |  | 


