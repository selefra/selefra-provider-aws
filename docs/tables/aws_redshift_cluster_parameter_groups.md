# Table: aws_redshift_cluster_parameter_groups

## Primary Keys 

```
cluster_arn, parameter_group_name
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| parameter_apply_status | string | X | √ |  | 
| aws_redshift_clusters_selefra_id | string | X | X | fk to aws_redshift_clusters.selefra_id | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| cluster_arn | string | X | √ | `The Amazon Resource Name (ARN) for the resource.` | 
| parameter_group_name | string | X | √ |  | 
| cluster_parameter_status_list | json | X | √ |  | 


