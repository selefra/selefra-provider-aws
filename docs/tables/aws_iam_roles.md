# Table: aws_iam_roles

## Primary Keys 

```
account_id, id
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| account_id | string | X | √ |  | 
| policies | json | X | √ |  | 
| id | string | X | √ |  | 
| tags | json | X | √ |  | 
| path | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| description | string | X | √ |  | 
| max_session_duration | int | X | √ |  | 
| permissions_boundary | json | X | √ |  | 
| assume_role_policy_document | json | X | √ |  | 
| arn | string | X | √ |  | 
| role_name | string | X | √ |  | 
| create_date | timestamp | X | √ |  | 
| role_last_used | json | X | √ |  | 


