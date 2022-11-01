# Table: aws_dax_clusters

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| tags | json | X | √ |  | 
| active_nodes | int | X | √ |  | 
| cluster_endpoint_encryption_type | string | X | √ |  | 
| cluster_name | string | X | √ |  | 
| preferred_maintenance_window | string | X | √ |  | 
| account_id | string | X | √ |  | 
| description | string | X | √ |  | 
| iam_role_arn | string | X | √ |  | 
| node_type | string | X | √ |  | 
| notification_configuration | json | X | √ |  | 
| sse_description | json | X | √ |  | 
| status | string | X | √ |  | 
| region | string | X | √ |  | 
| nodes | json | X | √ |  | 
| parameter_group | json | X | √ |  | 
| security_groups | json | X | √ |  | 
| subnet_group | string | X | √ |  | 
| arn | string | √ | √ |  | 
| cluster_discovery_endpoint | json | X | √ |  | 
| node_ids_to_remove | string_array | X | √ |  | 
| total_nodes | int | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 


