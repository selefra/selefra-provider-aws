# Table: aws_backup_vaults

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| notifications | json | X | √ |  | 
| creator_request_id | string | X | √ |  | 
| encryption_key_arn | string | X | √ |  | 
| locked | bool | X | √ |  | 
| number_of_recovery_points | int | X | √ |  | 
| region | string | X | √ |  | 
| arn | string | √ | √ |  | 
| creation_date | timestamp | X | √ |  | 
| min_retention_days | int | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| access_policy | json | X | √ |  | 
| tags | json | X | √ |  | 
| lock_date | timestamp | X | √ |  | 
| account_id | string | X | √ |  | 
| backup_vault_name | string | X | √ |  | 
| max_retention_days | int | X | √ |  | 


