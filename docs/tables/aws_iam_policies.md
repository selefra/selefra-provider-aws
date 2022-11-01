# Table: aws_iam_policies

## Primary Keys 

```
account_id, id
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| arn | string | X | √ |  | 
| attachment_count | int | X | √ |  | 
| create_date | timestamp | X | √ |  | 
| default_version_id | string | X | √ |  | 
| policy_name | string | X | √ |  | 
| account_id | string | X | √ |  | 
| id | string | X | √ |  | 
| policy_version_list | json | X | √ |  | 
| is_attachable | bool | X | √ |  | 
| update_date | timestamp | X | √ |  | 
| tags | json | X | √ |  | 
| permissions_boundary_usage_count | int | X | √ |  | 
| description | string | X | √ |  | 
| path | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 


