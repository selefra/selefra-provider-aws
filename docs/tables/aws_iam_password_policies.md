# Table: aws_iam_password_policies

## Primary Keys 

```
account_id
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| require_uppercase_characters | bool | X | √ |  | 
| allow_users_to_change_password | bool | X | √ |  | 
| hard_expiry | bool | X | √ |  | 
| max_password_age | int | X | √ |  | 
| password_reuse_prevention | int | X | √ |  | 
| require_lowercase_characters | bool | X | √ |  | 
| require_numbers | bool | X | √ |  | 
| require_symbols | bool | X | √ |  | 
| account_id | string | √ | √ |  | 
| expire_passwords | bool | X | √ |  | 
| minimum_password_length | int | X | √ |  | 
| policy_exists | bool | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 


