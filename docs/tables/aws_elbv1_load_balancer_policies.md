# Table: aws_elbv1_load_balancer_policies

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| load_balancer_arn | string | X | √ |  | 
| load_balancer_name | string | X | √ |  | 
| policy_type_name | string | X | √ |  | 
| aws_elbv1_load_balancers_selefra_id | string | X | X | fk to aws_elbv1_load_balancers.selefra_id | 
| selefra_id | string | √ | √ | random id | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| policy_attribute_descriptions | json | X | √ |  | 
| policy_name | string | X | √ |  | 


