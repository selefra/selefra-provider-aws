# Table: aws_redshift_subnet_groups

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| tags | json | X | √ | `The list of tags for the cluster subnet group.` | 
| description | string | X | √ |  | 
| subnets | json | X | √ |  | 
| region | string | X | √ |  | 
| arn | string | √ | √ | `The Amazon Resource Name (ARN) for the resource.` | 
| subnet_group_status | string | X | √ |  | 
| vpc_id | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| account_id | string | X | √ |  | 
| cluster_subnet_group_name | string | X | √ |  | 


