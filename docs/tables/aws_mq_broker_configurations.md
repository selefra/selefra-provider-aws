# Table: aws_mq_broker_configurations

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| account_id | string | X | √ |  | 
| broker_arn | string | X | √ |  | 
| engine_type | string | X | √ |  | 
| selefra_id | string | √ | √ | random id | 
| region | string | X | √ |  | 
| authentication_strategy | string | X | √ |  | 
| created | timestamp | X | √ |  | 
| id | string | X | √ |  | 
| name | string | X | √ |  | 
| aws_mq_brokers_selefra_id | string | X | X | fk to aws_mq_brokers.selefra_id | 
| description | string | X | √ |  | 
| latest_revision | json | X | √ |  | 
| tags | json | X | √ |  | 
| arn | string | X | √ |  | 
| engine_version | string | X | √ |  | 


