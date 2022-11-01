# Table: aws_eventbridge_event_bus_rules

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| state | string | X | √ |  | 
| account_id | string | X | √ |  | 
| name | string | X | √ |  | 
| tags | json | X | √ |  | 
| description | string | X | √ |  | 
| event_bus_name | string | X | √ |  | 
| managed_by | string | X | √ |  | 
| schedule_expression | string | X | √ |  | 
| region | string | X | √ |  | 
| event_bus_arn | string | X | √ |  | 
| aws_eventbridge_event_buses_selefra_id | string | X | X | fk to aws_eventbridge_event_buses.selefra_id | 
| role_arn | string | X | √ |  | 
| selefra_id | string | √ | √ | random id | 
| arn | string | X | √ |  | 
| event_pattern | string | X | √ |  | 


