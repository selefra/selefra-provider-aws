# Table: aws_iam_server_certificates

## Primary Keys 

```
account_id, id
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| path | string | X | √ |  | 
| server_certificate_name | string | X | √ |  | 
| expiration | timestamp | X | √ |  | 
| upload_date | timestamp | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| account_id | string | X | √ |  | 
| id | string | X | √ |  | 
| arn | string | X | √ |  | 


