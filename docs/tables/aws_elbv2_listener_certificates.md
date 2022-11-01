# Table: aws_elbv2_listener_certificates

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| listener_arn | string | X | √ |  | 
| certificate_arn | string | X | √ |  | 
| is_default | bool | X | √ |  | 
| selefra_id | string | √ | √ | random id | 
| aws_elbv2_listeners_selefra_id | string | X | X | fk to aws_elbv2_listeners.selefra_id | 


