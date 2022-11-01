# Table: aws_apigatewayv2_api_integrations

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| arn | string | X | √ |  | 
| integration_id | string | X | √ |  | 
| integration_response_selection_expression | string | X | √ |  | 
| tls_config | json | X | √ |  | 
| aws_apigatewayv2_apis_selefra_id | string | X | X | fk to aws_apigatewayv2_apis.selefra_id | 
| request_templates | json | X | √ |  | 
| connection_id | string | X | √ |  | 
| integration_type | string | X | √ |  | 
| response_parameters | json | X | √ |  | 
| selefra_id | string | √ | √ | random id | 
| content_handling_strategy | string | X | √ |  | 
| credentials_arn | string | X | √ |  | 
| integration_method | string | X | √ |  | 
| integration_subtype | string | X | √ |  | 
| integration_uri | string | X | √ |  | 
| payload_format_version | string | X | √ |  | 
| account_id | string | X | √ |  | 
| description | string | X | √ |  | 
| passthrough_behavior | string | X | √ |  | 
| region | string | X | √ |  | 
| api_gateway_managed | bool | X | √ |  | 
| connection_type | string | X | √ |  | 
| api_arn | string | X | √ |  | 
| api_id | string | X | √ |  | 
| request_parameters | json | X | √ |  | 
| template_selection_expression | string | X | √ |  | 
| timeout_in_millis | int | X | √ |  | 


