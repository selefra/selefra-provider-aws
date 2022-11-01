# Table: aws_iot_streams

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| stream_id | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| account_id | string | X | √ |  | 
| description | string | X | √ |  | 
| files | json | X | √ |  | 
| last_updated_at | timestamp | X | √ |  | 
| stream_version | int | X | √ |  | 
| region | string | X | √ |  | 
| arn | string | √ | √ |  | 
| created_at | timestamp | X | √ |  | 
| role_arn | string | X | √ |  | 


