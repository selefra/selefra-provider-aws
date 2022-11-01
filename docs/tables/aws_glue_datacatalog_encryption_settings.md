# Table: aws_glue_datacatalog_encryption_settings

## Primary Keys 

```
account_id
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| account_id | string | √ | √ |  | 
| region | string | X | √ |  | 
| connection_password_encryption | json | X | √ |  | 
| encryption_at_rest | json | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 


