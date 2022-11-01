# Table: aws_apigateway_domain_name_base_path_mappings

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| arn | string | X | √ |  | 
| base_path | string | X | √ |  | 
| rest_api_id | string | X | √ |  | 
| aws_apigateway_domain_names_selefra_id | string | X | X | fk to aws_apigateway_domain_names.selefra_id | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| domain_name_arn | string | X | √ |  | 
| stage | string | X | √ |  | 
| selefra_id | string | √ | √ | random id | 


