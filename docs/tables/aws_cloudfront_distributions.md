# Table: aws_cloudfront_distributions

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| active_trusted_signers | json | X | √ |  | 
| alias_icp_recordals | json | X | √ |  | 
| account_id | string | X | √ |  | 
| domain_name | string | X | √ |  | 
| in_progress_invalidation_batches | int | X | √ |  | 
| last_modified_time | timestamp | X | √ |  | 
| status | string | X | √ |  | 
| active_trusted_key_groups | json | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| tags | json | X | √ |  | 
| arn | string | √ | √ |  | 
| distribution_config | json | X | √ |  | 
| id | string | X | √ |  | 


