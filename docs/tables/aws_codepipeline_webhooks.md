# Table: aws_codepipeline_webhooks

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| arn | string | √ | √ |  | 
| definition | json | X | √ |  | 
| url | string | X | √ |  | 
| error_code | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| error_message | string | X | √ |  | 
| last_triggered | timestamp | X | √ |  | 
| tags | json | X | √ |  | 


