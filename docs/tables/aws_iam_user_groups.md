# Table: aws_iam_user_groups

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| group_id | string | X | √ |  | 
| group_name | string | X | √ |  | 
| aws_iam_users_selefra_id | string | X | X | fk to aws_iam_users.selefra_id | 
| account_id | string | X | √ |  | 
| user_arn | string | X | √ |  | 
| arn | string | X | √ |  | 
| create_date | timestamp | X | √ |  | 
| user_id | string | X | √ |  | 
| path | string | X | √ |  | 
| selefra_id | string | √ | √ | random id | 


