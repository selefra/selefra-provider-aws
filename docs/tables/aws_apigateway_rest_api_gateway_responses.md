# Table: aws_apigateway_rest_api_gateway_responses

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| response_type | string | X | √ |  | 
| status_code | string | X | √ |  | 
| region | string | X | √ |  | 
| rest_api_arn | string | X | √ |  | 
| arn | string | X | √ |  | 
| default_response | bool | X | √ |  | 
| response_parameters | json | X | √ |  | 
| response_templates | json | X | √ |  | 
| selefra_id | string | √ | √ | random id | 
| account_id | string | X | √ |  | 
| aws_apigateway_rest_apis_selefra_id | string | X | X | fk to aws_apigateway_rest_apis.selefra_id | 


