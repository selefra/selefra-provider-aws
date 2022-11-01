# Table: aws_lightsail_bucket_access_keys

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| account_id | string | X | √ |  | 
| bucket_arn | string | X | √ |  | 
| created_at | timestamp | X | √ |  | 
| secret_access_key | string | X | √ |  | 
| status | string | X | √ |  | 
| selefra_id | string | √ | √ | random id | 
| region | string | X | √ |  | 
| access_key_id | string | X | √ |  | 
| last_used | json | X | √ |  | 
| aws_lightsail_buckets_selefra_id | string | X | X | fk to aws_lightsail_buckets.selefra_id | 


