# Table: aws_apigateway_rest_apis

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| api_key_source | string | X | √ |  | 
| binary_media_types | string_array | X | √ |  | 
| tags | json | X | √ |  | 
| version | string | X | √ |  | 
| warnings | string_array | X | √ |  | 
| account_id | string | X | √ |  | 
| arn | string | √ | √ |  | 
| policy | string | X | √ |  | 
| id | string | X | √ |  | 
| minimum_compression_size | int | X | √ |  | 
| disable_execute_api_endpoint | bool | X | √ |  | 
| endpoint_configuration | json | X | √ |  | 
| region | string | X | √ |  | 
| description | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| created_date | timestamp | X | √ |  | 
| name | string | X | √ |  | 


