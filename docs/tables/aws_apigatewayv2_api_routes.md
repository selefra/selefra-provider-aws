# Table: aws_apigatewayv2_api_routes

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| region | string | X | √ |  | 
| route_id | string | X | √ |  | 
| api_id | string | X | √ |  | 
| route_key | string | X | √ |  | 
| api_key_required | bool | X | √ |  | 
| model_selection_expression | string | X | √ |  | 
| request_parameters | json | X | √ |  | 
| route_response_selection_expression | string | X | √ |  | 
| account_id | string | X | √ |  | 
| api_arn | string | X | √ |  | 
| authorization_scopes | string_array | X | √ |  | 
| target | string | X | √ |  | 
| arn | string | X | √ |  | 
| api_gateway_managed | bool | X | √ |  | 
| authorization_type | string | X | √ |  | 
| authorizer_id | string | X | √ |  | 
| operation_name | string | X | √ |  | 
| request_models | json | X | √ |  | 
| aws_apigatewayv2_apis_selefra_id | string | X | X | fk to aws_apigatewayv2_apis.selefra_id | 
| selefra_id | string | √ | √ | random id | 


