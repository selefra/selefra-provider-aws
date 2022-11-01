# Table: aws_appsync_graphql_apis

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| user_pool_config | json | X | √ |  | 
| waf_web_acl_arn | string | X | √ |  | 
| xray_enabled | bool | X | √ |  | 
| name | string | X | √ |  | 
| open_id_connect_config | json | X | √ |  | 
| region | string | X | √ |  | 
| arn | string | √ | √ |  | 
| additional_authentication_providers | json | X | √ |  | 
| api_id | string | X | √ |  | 
| authentication_type | string | X | √ |  | 
| log_config | json | X | √ |  | 
| tags | json | X | √ |  | 
| uris | json | X | √ |  | 
| account_id | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| lambda_authorizer_config | json | X | √ |  | 


