# Table: aws_cognito_identity_pools

## Primary Keys 

```
account_id, region, id
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| arn | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| identity_pool_name | string | X | √ |  | 
| allow_classic_flow | bool | X | √ |  | 
| cognito_identity_providers | json | X | √ |  | 
| identity_pool_tags | json | X | √ |  | 
| open_id_connect_provider_ar_ns | string_array | X | √ |  | 
| account_id | string | X | √ |  | 
| result_metadata | json | X | √ |  | 
| region | string | X | √ |  | 
| id | string | X | √ |  | 
| allow_unauthenticated_identities | bool | X | √ |  | 
| developer_provider_name | string | X | √ |  | 
| saml_provider_ar_ns | string_array | X | √ |  | 
| supported_login_providers | json | X | √ |  | 


