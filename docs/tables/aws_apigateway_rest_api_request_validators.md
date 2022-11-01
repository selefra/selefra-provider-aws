# Table: aws_apigateway_rest_api_request_validators

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| rest_api_arn | string | X | √ |  | 
| arn | string | X | √ |  | 
| id | string | X | √ |  | 
| name | string | X | √ |  | 
| account_id | string | X | √ |  | 
| validate_request_body | bool | X | √ |  | 
| validate_request_parameters | bool | X | √ |  | 
| aws_apigateway_rest_apis_selefra_id | string | X | X | fk to aws_apigateway_rest_apis.selefra_id | 
| selefra_id | string | √ | √ | random id | 
| region | string | X | √ |  | 


