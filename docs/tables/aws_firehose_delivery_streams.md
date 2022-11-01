# Table: aws_firehose_delivery_streams

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| region | string | X | √ |  | 
| arn | string | √ | √ |  | 
| delivery_stream_status | string | X | √ |  | 
| last_update_timestamp | timestamp | X | √ |  | 
| account_id | string | X | √ |  | 
| tags | json | X | √ |  | 
| has_more_destinations | bool | X | √ |  | 
| delivery_stream_encryption_configuration | json | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| version_id | string | X | √ |  | 
| source | json | X | √ |  | 
| delivery_stream_name | string | X | √ |  | 
| delivery_stream_type | string | X | √ |  | 
| destinations | json | X | √ |  | 
| create_timestamp | timestamp | X | √ |  | 
| failure_description | json | X | √ |  | 


