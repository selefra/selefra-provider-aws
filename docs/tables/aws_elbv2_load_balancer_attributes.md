# Table: aws_elbv2_load_balancer_attributes

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| load_balancer_arn | string | X | √ |  | 
| key | string | X | √ |  | 
| value | string | X | √ |  | 
| aws_elbv2_load_balancers_selefra_id | string | X | X | fk to aws_elbv2_load_balancers.selefra_id | 
| selefra_id | string | √ | √ | random id | 


