# Table: aws_apigateway_rest_api_stages

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| last_updated_date | timestamp | X | √ |  | 
| stage_name | string | X | √ |  | 
| arn | string | X | √ |  | 
| canary_settings | json | X | √ |  | 
| client_certificate_id | string | X | √ |  | 
| created_date | timestamp | X | √ |  | 
| description | string | X | √ |  | 
| tracing_enabled | bool | X | √ |  | 
| rest_api_arn | string | X | √ |  | 
| deployment_id | string | X | √ |  | 
| documentation_version | string | X | √ |  | 
| variables | json | X | √ |  | 
| web_acl_arn | string | X | √ |  | 
| selefra_id | string | √ | √ | random id | 
| method_settings | json | X | √ |  | 
| tags | json | X | √ |  | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| access_log_settings | json | X | √ |  | 
| cache_cluster_enabled | bool | X | √ |  | 
| cache_cluster_size | string | X | √ |  | 
| cache_cluster_status | string | X | √ |  | 
| aws_apigateway_rest_apis_selefra_id | string | X | X | fk to aws_apigateway_rest_apis.selefra_id | 


