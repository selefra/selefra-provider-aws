# Table: aws_kinesis_streams

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| tags | json | X | √ |  | 
| stream_creation_timestamp | timestamp | X | √ |  | 
| consumer_count | int | X | √ |  | 
| stream_mode_details | json | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| arn | string | √ | √ |  | 
| region | string | X | √ |  | 
| enhanced_monitoring | json | X | √ |  | 
| encryption_type | string | X | √ |  | 
| account_id | string | X | √ |  | 
| retention_period_hours | int | X | √ |  | 
| open_shard_count | int | X | √ |  | 
| stream_status | string | X | √ |  | 
| key_id | string | X | √ |  | 
| stream_name | string | X | √ |  | 


