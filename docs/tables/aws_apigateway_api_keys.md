# Table: aws_apigateway_api_keys

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| account_id | string | X | √ |  | 
| created_date | timestamp | X | √ |  | 
| customer_id | string | X | √ |  | 
| name | string | X | √ |  | 
| stage_keys | string_array | X | √ |  | 
| description | string | X | √ |  | 
| last_updated_date | timestamp | X | √ |  | 
| tags | json | X | √ |  | 
| value | string | X | √ |  | 
| region | string | X | √ |  | 
| arn | string | √ | √ |  | 
| enabled | bool | X | √ |  | 
| id | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 


