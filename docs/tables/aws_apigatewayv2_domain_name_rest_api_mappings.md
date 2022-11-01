# Table: aws_apigatewayv2_domain_name_rest_api_mappings

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| selefra_id | string | √ | √ | random id | 
| aws_apigatewayv2_domain_names_selefra_id | string | X | X | fk to aws_apigatewayv2_domain_names.selefra_id | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| arn | string | X | √ |  | 
| api_id | string | X | √ |  | 
| api_mapping_id | string | X | √ |  | 
| api_mapping_key | string | X | √ |  | 
| domain_name_arn | string | X | √ |  | 
| stage | string | X | √ |  | 


