# Table: aws_iam_openid_connect_identity_providers

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| client_id_list | string_array | X | √ |  | 
| thumbprint_list | string_array | X | √ |  | 
| url | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| account_id | string | X | √ |  | 
| arn | string | √ | √ |  | 
| result_metadata | json | X | √ |  | 
| tags | json | X | √ |  | 
| create_date | timestamp | X | √ |  | 


