# Table: aws_iam_users

## Primary Keys 

```
id, account_id
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| path | string | X | √ |  | 
| permissions_boundary | json | X | √ |  | 
| arn | string | X | √ |  | 
| account_id | string | X | √ |  | 
| tags | json | X | √ |  | 
| create_date | timestamp | X | √ |  | 
| user_name | string | X | √ |  | 
| password_last_used | timestamp | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| id | string | X | √ |  | 


