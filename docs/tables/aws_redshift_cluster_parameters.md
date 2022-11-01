# Table: aws_redshift_cluster_parameters

## Primary Keys 

```
cluster_arn, parameter_name
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| apply_type | string | X | √ |  | 
| source | string | X | √ |  | 
| parameter_name | string | X | √ |  | 
| description | string | X | √ |  | 
| is_modifiable | bool | X | √ |  | 
| minimum_engine_version | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| region | string | X | √ |  | 
| cluster_arn | string | X | √ | `The Amazon Resource Name (ARN) for the resource.` | 
| allowed_values | string | X | √ |  | 
| account_id | string | X | √ |  | 
| data_type | string | X | √ |  | 
| parameter_value | string | X | √ |  | 
| aws_redshift_cluster_parameter_groups_selefra_id | string | X | X | fk to aws_redshift_cluster_parameter_groups.selefra_id | 


