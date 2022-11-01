# Table: aws_s3_buckets

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| creation_date | timestamp | X | √ |  | 
| name | string | X | √ |  | 
| block_public_policy | bool | X | √ |  | 
| ownership_controls | string_array | X | √ |  | 
| region | string | X | √ |  | 
| logging_target_bucket | string | X | √ |  | 
| versioning_status | string | X | √ |  | 
| tags | json | X | √ |  | 
| arn | string | √ | √ |  | 
| replication_role | string | X | √ |  | 
| logging_target_prefix | string | X | √ |  | 
| block_public_acls | bool | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| account_id | string | X | √ |  | 
| replication_rules | json | X | √ |  | 
| policy | json | X | √ |  | 
| versioning_mfa_delete | string | X | √ |  | 
| ignore_public_acls | bool | X | √ |  | 
| restrict_public_buckets | bool | X | √ |  | 


