# Table: aws_mq_broker_users

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| account_id | string | X | √ |  | 
| broker_arn | string | X | √ |  | 
| console_access | bool | X | √ |  | 
| username | string | X | √ |  | 
| aws_mq_brokers_selefra_id | string | X | X | fk to aws_mq_brokers.selefra_id | 
| region | string | X | √ |  | 
| broker_id | string | X | √ |  | 
| groups | string_array | X | √ |  | 
| pending | json | X | √ |  | 
| result_metadata | json | X | √ |  | 
| selefra_id | string | √ | √ | random id | 


