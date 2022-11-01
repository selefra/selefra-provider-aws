# Table: aws_apigateway_rest_api_deployments

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| description | string | X | √ |  | 
| aws_apigateway_rest_apis_selefra_id | string | X | X | fk to aws_apigateway_rest_apis.selefra_id | 
| account_id | string | X | √ |  | 
| rest_api_arn | string | X | √ |  | 
| arn | string | X | √ |  | 
| api_summary | json | X | √ |  | 
| created_date | timestamp | X | √ |  | 
| id | string | X | √ |  | 
| selefra_id | string | √ | √ | random id | 
| region | string | X | √ |  | 


