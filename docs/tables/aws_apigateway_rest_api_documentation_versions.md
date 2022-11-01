# Table: aws_apigateway_rest_api_documentation_versions

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| arn | string | X | √ |  | 
| created_date | timestamp | X | √ |  | 
| account_id | string | X | √ |  | 
| rest_api_arn | string | X | √ |  | 
| version | string | X | √ |  | 
| aws_apigateway_rest_apis_selefra_id | string | X | X | fk to aws_apigateway_rest_apis.selefra_id | 
| selefra_id | string | √ | √ | random id | 
| region | string | X | √ |  | 
| description | string | X | √ |  | 


