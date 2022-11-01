# Table: aws_sqs_queues

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| url | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| content_based_deduplication | bool | X | √ |  | 
| approximate_number_of_messages_delayed | int | X | √ |  | 
| approximate_number_of_messages_not_visible | int | X | √ |  | 
| sqs_managed_sse_enabled | bool | X | √ |  | 
| redrive_policy | json | X | √ |  | 
| visibility_timeout | int | X | √ |  | 
| fifo_queue | bool | X | √ |  | 
| delay_seconds | int | X | √ |  | 
| last_modified_timestamp | int | X | √ |  | 
| message_retention_period | int | X | √ |  | 
| fifo_throughput_limit | string | X | √ |  | 
| region | string | X | √ |  | 
| arn | string | √ | √ |  | 
| approximate_number_of_messages | int | X | √ |  | 
| kms_data_key_reuse_period_seconds | int | X | √ |  | 
| deduplication_scope | string | X | √ |  | 
| tags | json | X | √ |  | 
| maximum_message_size | int | X | √ |  | 
| redrive_allow_policy | json | X | √ |  | 
| kms_master_key_id | string | X | √ |  | 
| unknown_fields | json | X | √ |  | 
| account_id | string | X | √ |  | 
| policy | json | X | √ |  | 
| created_timestamp | int | X | √ |  | 
| receive_message_wait_time_seconds | int | X | √ |  | 


