# Table: aws_sns_subscriptions

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| arn | string | √ | √ |  | 
| delivery_policy | json | X | √ |  | 
| redrive_policy | json | X | √ |  | 
| owner | string | X | √ |  | 
| effective_delivery_policy | json | X | √ |  | 
| protocol | string | X | √ |  | 
| confirmation_was_authenticated | bool | X | √ |  | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| topic_arn | string | X | √ |  | 
| raw_message_delivery | bool | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| filter_policy | json | X | √ |  | 
| endpoint | string | X | √ |  | 
| pending_confirmation | bool | X | √ |  | 
| subscription_role_arn | string | X | √ |  | 
| unknown_fields | json | X | √ |  | 


