# Table: aws_apigateway_usage_plan_keys

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| id | string | X | √ |  | 
| name | string | X | √ |  | 
| type | string | X | √ |  | 
| value | string | X | √ |  | 
| selefra_id | string | √ | √ | random id | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| aws_apigateway_usage_plans_selefra_id | string | X | X | fk to aws_apigateway_usage_plans.selefra_id | 
| usage_plan_arn | string | X | √ |  | 
| arn | string | X | √ |  | 


