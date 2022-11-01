# Table: aws_shield_subscriptions

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| time_commitment_in_seconds | int | X | √ |  | 
| account_id | string | X | √ |  | 
| subscription_limits | json | X | √ |  | 
| end_time | timestamp | X | √ |  | 
| proactive_engagement_status | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| arn | string | √ | √ |  | 
| auto_renew | string | X | √ |  | 
| limits | json | X | √ |  | 
| start_time | timestamp | X | √ |  | 


