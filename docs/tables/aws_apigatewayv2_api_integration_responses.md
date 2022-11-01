# Table: aws_apigatewayv2_api_integration_responses

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| response_parameters | json | X | √ |  | 
| response_templates | json | X | √ |  | 
| region | string | X | √ |  | 
| api_integration_arn | string | X | √ |  | 
| arn | string | X | √ |  | 
| content_handling_strategy | string | X | √ |  | 
| integration_response_id | string | X | √ |  | 
| template_selection_expression | string | X | √ |  | 
| aws_apigatewayv2_api_integrations_selefra_id | string | X | X | fk to aws_apigatewayv2_api_integrations.selefra_id | 
| selefra_id | string | √ | √ | random id | 
| account_id | string | X | √ |  | 
| integration_id | string | X | √ |  | 
| integration_response_key | string | X | √ |  | 


