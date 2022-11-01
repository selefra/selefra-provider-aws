# Table: aws_apigatewayv2_api_deployments

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| api_id | string | X | √ |  | 
| deployment_status_message | string | X | √ |  | 
| selefra_id | string | √ | √ | random id | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| api_arn | string | X | √ |  | 
| arn | string | X | √ |  | 
| auto_deployed | bool | X | √ |  | 
| created_date | timestamp | X | √ |  | 
| deployment_id | string | X | √ |  | 
| deployment_status | string | X | √ |  | 
| description | string | X | √ |  | 
| aws_apigatewayv2_apis_selefra_id | string | X | X | fk to aws_apigatewayv2_apis.selefra_id | 


