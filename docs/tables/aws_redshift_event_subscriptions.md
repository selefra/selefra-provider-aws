# Table: aws_redshift_event_subscriptions

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| source_ids_list | string_array | X | √ |  | 
| subscription_creation_time | timestamp | X | √ |  | 
| account_id | string | X | √ |  | 
| severity | string | X | √ |  | 
| source_type | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| arn | string | √ | √ | `ARN of the event subscription.` | 
| tags | json | X | √ | `Tags` | 
| cust_subscription_id | string | X | √ |  | 
| customer_aws_id | string | X | √ |  | 
| enabled | bool | X | √ |  | 
| event_categories_list | string_array | X | √ |  | 
| status | string | X | √ |  | 
| region | string | X | √ |  | 
| sns_topic_arn | string | X | √ |  | 


