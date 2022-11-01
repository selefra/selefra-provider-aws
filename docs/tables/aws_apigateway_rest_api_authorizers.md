# Table: aws_apigateway_rest_api_authorizers

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| region | string | X | √ |  | 
| rest_api_arn | string | X | √ |  | 
| authorizer_credentials | string | X | √ |  | 
| authorizer_result_ttl_in_seconds | int | X | √ |  | 
| identity_validation_expression | string | X | √ |  | 
| aws_apigateway_rest_apis_selefra_id | string | X | X | fk to aws_apigateway_rest_apis.selefra_id | 
| account_id | string | X | √ |  | 
| provider_ar_ns | string_array | X | √ |  | 
| type | string | X | √ |  | 
| auth_type | string | X | √ |  | 
| id | string | X | √ |  | 
| name | string | X | √ |  | 
| arn | string | X | √ |  | 
| authorizer_uri | string | X | √ |  | 
| identity_source | string | X | √ |  | 
| selefra_id | string | √ | √ | random id | 


