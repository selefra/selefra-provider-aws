# Table: aws_backup_vault_recovery_points

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| resource_arn | string | X | √ |  | 
| vault_arn | string | X | √ |  | 
| iam_role_arn | string | X | √ |  | 
| arn | string | √ | √ |  | 
| backup_vault_arn | string | X | √ |  | 
| creation_date | timestamp | X | √ |  | 
| last_restore_time | timestamp | X | √ |  | 
| status_message | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| calculated_lifecycle | json | X | √ |  | 
| created_by | json | X | √ |  | 
| source_backup_vault_arn | string | X | √ |  | 
| status | string | X | √ |  | 
| tags | json | X | √ |  | 
| backup_size_in_bytes | int | X | √ |  | 
| encryption_key_arn | string | X | √ |  | 
| is_encrypted | bool | X | √ |  | 
| lifecycle | json | X | √ |  | 
| resource_type | string | X | √ |  | 
| aws_backup_vaults_selefra_id | string | X | X | fk to aws_backup_vaults.selefra_id | 
| backup_vault_name | string | X | √ |  | 
| completion_date | timestamp | X | √ |  | 


