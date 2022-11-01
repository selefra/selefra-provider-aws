# Table: aws_apigatewayv2_api_authorizers

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| enable_simple_responses | bool | X | √ |  | 
| account_id | string | X | √ |  | 
| authorizer_payload_format_version | string | X | √ |  | 
| authorizer_result_ttl_in_seconds | int | X | √ |  | 
| authorizer_type | string | X | √ |  | 
| api_arn | string | X | √ |  | 
| arn | string | X | √ |  | 
| selefra_id | string | √ | √ | random id | 
| authorizer_uri | string | X | √ |  | 
| identity_source | string_array | X | √ |  | 
| identity_validation_expression | string | X | √ |  | 
| region | string | X | √ |  | 
| api_id | string | X | √ |  | 
| name | string | X | √ |  | 
| authorizer_credentials_arn | string | X | √ |  | 
| authorizer_id | string | X | √ |  | 
| jwt_configuration | json | X | √ |  | 
| aws_apigatewayv2_apis_selefra_id | string | X | X | fk to aws_apigatewayv2_apis.selefra_id | 


