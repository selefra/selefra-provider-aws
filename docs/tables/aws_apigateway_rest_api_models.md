# Table: aws_apigateway_rest_api_models

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| selefra_id | string | √ | √ | random id | 
| rest_api_arn | string | X | √ |  | 
| arn | string | X | √ |  | 
| content_type | string | X | √ |  | 
| description | string | X | √ |  | 
| aws_apigateway_rest_apis_selefra_id | string | X | X | fk to aws_apigateway_rest_apis.selefra_id | 
| schema | string | X | √ |  | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| model_template | string | X | √ |  | 
| id | string | X | √ |  | 
| name | string | X | √ |  | 


