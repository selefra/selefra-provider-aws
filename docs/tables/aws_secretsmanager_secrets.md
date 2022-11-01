# Table: aws_secretsmanager_secrets

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| last_accessed_date | timestamp | X | √ |  | 
| last_changed_date | timestamp | X | √ |  | 
| name | string | X | √ |  | 
| rotation_lambda_arn | string | X | √ |  | 
| rotation_rules | json | X | √ |  | 
| account_id | string | X | √ |  | 
| tags | json | X | √ | `The list of user-defined tags associated with the secret` | 
| description | string | X | √ |  | 
| primary_region | string | X | √ |  | 
| rotation_enabled | bool | X | √ |  | 
| replication_status | json | X | √ |  | 
| version_ids_to_stages | json | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| arn | string | √ | √ |  | 
| policy | json | X | √ | `A JSON-formatted string that describes the permissions that are associated with the attached secret.` | 
| deleted_date | timestamp | X | √ |  | 
| last_rotated_date | timestamp | X | √ |  | 
| owning_service | string | X | √ |  | 
| region | string | X | √ |  | 
| created_date | timestamp | X | √ |  | 
| kms_key_id | string | X | √ |  | 


