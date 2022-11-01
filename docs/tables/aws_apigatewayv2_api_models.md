# Table: aws_apigatewayv2_api_models

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| model_template | string | X | √ |  | 
| content_type | string | X | √ |  | 
| description | string | X | √ |  | 
| model_id | string | X | √ |  | 
| schema | string | X | √ |  | 
| selefra_id | string | √ | √ | random id | 
| region | string | X | √ |  | 
| api_arn | string | X | √ |  | 
| aws_apigatewayv2_apis_selefra_id | string | X | X | fk to aws_apigatewayv2_apis.selefra_id | 
| arn | string | X | √ |  | 
| name | string | X | √ |  | 
| account_id | string | X | √ |  | 
| api_id | string | X | √ |  | 


