# Table: aws_sns_topics

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| arn | string | √ | √ |  | 
| subscriptions_pending | int | X | √ |  | 
| fifo_topic | bool | X | √ |  | 
| content_based_deduplication | bool | X | √ |  | 
| region | string | X | √ |  | 
| effective_delivery_policy | json | X | √ |  | 
| display_name | string | X | √ |  | 
| subscriptions_confirmed | int | X | √ |  | 
| tags | json | X | √ |  | 
| subscriptions_deleted | int | X | √ |  | 
| kms_master_key_id | string | X | √ |  | 
| unknown_fields | json | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| account_id | string | X | √ |  | 
| policy | json | X | √ |  | 
| owner | string | X | √ |  | 
| delivery_policy | json | X | √ |  | 


