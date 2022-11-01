# Table: aws_s3_bucket_grants

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| bucket_arn | string | X | √ |  | 
| grantee | json | X | √ |  | 
| permission | string | X | √ |  | 
| aws_s3_buckets_selefra_id | string | X | X | fk to aws_s3_buckets.selefra_id | 
| selefra_id | string | √ | √ | random id | 
| account_id | string | X | √ |  | 


