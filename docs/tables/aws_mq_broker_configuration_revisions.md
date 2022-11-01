# Table: aws_mq_broker_configuration_revisions

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| broker_configuration_arn | string | X | √ |  | 
| data | json | X | √ |  | 
| configuration_id | string | X | √ |  | 
| description | string | X | √ |  | 
| result_metadata | json | X | √ |  | 
| aws_mq_broker_configurations_selefra_id | string | X | X | fk to aws_mq_broker_configurations.selefra_id | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| created | timestamp | X | √ |  | 
| selefra_id | string | √ | √ | random id | 


