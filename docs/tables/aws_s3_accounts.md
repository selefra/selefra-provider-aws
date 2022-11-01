# Table: aws_s3_accounts

## Primary Keys 

```
account_id
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| config_exists | bool | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| account_id | string | √ | √ |  | 
| block_public_acls | bool | X | √ |  | 
| block_public_policy | bool | X | √ |  | 
| ignore_public_acls | bool | X | √ |  | 
| restrict_public_buckets | bool | X | √ |  | 


