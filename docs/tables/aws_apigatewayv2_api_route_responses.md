# Table: aws_apigatewayv2_api_route_responses

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| model_selection_expression | string | X | √ |  | 
| response_models | json | X | √ |  | 
| region | string | X | √ |  | 
| route_id | string | X | √ |  | 
| route_response_key | string | X | √ |  | 
| response_parameters | json | X | √ |  | 
| route_response_id | string | X | √ |  | 
| selefra_id | string | √ | √ | random id | 
| aws_apigatewayv2_api_routes_selefra_id | string | X | X | fk to aws_apigatewayv2_api_routes.selefra_id | 
| account_id | string | X | √ |  | 
| api_route_arn | string | X | √ |  | 
| arn | string | X | √ |  | 


