# Table: aws_apigatewayv2_api_stages

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| api_arn | string | X | √ |  | 
| arn | string | X | √ |  | 
| access_log_settings | json | X | √ |  | 
| last_updated_date | timestamp | X | √ |  | 
| route_settings | json | X | √ |  | 
| aws_apigatewayv2_apis_selefra_id | string | X | X | fk to aws_apigatewayv2_apis.selefra_id | 
| account_id | string | X | √ |  | 
| last_deployment_status_message | string | X | √ |  | 
| tags | json | X | √ |  | 
| selefra_id | string | √ | √ | random id | 
| region | string | X | √ |  | 
| stage_name | string | X | √ |  | 
| auto_deploy | bool | X | √ |  | 
| client_certificate_id | string | X | √ |  | 
| default_route_settings | json | X | √ |  | 
| stage_variables | json | X | √ |  | 
| api_id | string | X | √ |  | 
| api_gateway_managed | bool | X | √ |  | 
| created_date | timestamp | X | √ |  | 
| deployment_id | string | X | √ |  | 
| description | string | X | √ |  | 


