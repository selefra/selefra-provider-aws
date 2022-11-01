# Table: aws_neptune_event_subscriptions

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| region | string | X | √ |  | 
| source_ids_list | string_array | X | √ |  | 
| account_id | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| arn | string | √ | √ |  | 
| status | string | X | √ |  | 
| subscription_creation_time | string | X | √ |  | 
| enabled | bool | X | √ |  | 
| event_categories_list | string_array | X | √ |  | 
| sns_topic_arn | string | X | √ |  | 
| source_type | string | X | √ |  | 
| tags | json | X | √ |  | 
| cust_subscription_id | string | X | √ |  | 
| customer_aws_id | string | X | √ |  | 


