# Table: aws_s3_bucket_cors_rules

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| allowed_origins | string_array | X | √ |  | 
| allowed_headers | string_array | X | √ |  | 
| expose_headers | string_array | X | √ |  | 
| max_age_seconds | int | X | √ |  | 
| aws_s3_buckets_selefra_id | string | X | X | fk to aws_s3_buckets.selefra_id | 
| account_id | string | X | √ |  | 
| bucket_arn | string | X | √ |  | 
| allowed_methods | string_array | X | √ |  | 
| id | string | X | √ |  | 
| selefra_id | string | √ | √ | random id | 


