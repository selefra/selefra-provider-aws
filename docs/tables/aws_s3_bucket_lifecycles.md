# Table: aws_s3_bucket_lifecycles

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| abort_incomplete_multipart_upload | json | X | √ |  | 
| noncurrent_version_expiration | json | X | √ |  | 
| noncurrent_version_transitions | json | X | √ |  | 
| prefix | string | X | √ |  | 
| aws_s3_buckets_selefra_id | string | X | X | fk to aws_s3_buckets.selefra_id | 
| selefra_id | string | √ | √ | random id | 
| status | string | X | √ |  | 
| bucket_arn | string | X | √ |  | 
| expiration | json | X | √ |  | 
| id | string | X | √ |  | 
| transitions | json | X | √ |  | 
| account_id | string | X | √ |  | 


