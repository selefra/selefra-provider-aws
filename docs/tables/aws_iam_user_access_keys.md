# Table: aws_iam_user_access_keys

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| user_id | string | X | √ |  | 
| create_date | timestamp | X | √ |  | 
| user_name | string | X | √ |  | 
| aws_iam_users_selefra_id | string | X | X | fk to aws_iam_users.selefra_id | 
| account_id | string | X | √ |  | 
| last_used | timestamp | X | √ |  | 
| last_used_service_name | string | X | √ |  | 
| access_key_id | string | X | √ |  | 
| status | string | X | √ |  | 
| last_rotated | timestamp | X | √ |  | 
| selefra_id | string | √ | √ | random id | 
| user_arn | string | X | √ |  | 


