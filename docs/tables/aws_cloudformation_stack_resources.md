# Table: aws_cloudformation_stack_resources

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| account_id | string | X | √ |  | 
| last_updated_timestamp | timestamp | X | √ |  | 
| logical_resource_id | string | X | √ |  | 
| resource_type | string | X | √ |  | 
| physical_resource_id | string | X | √ |  | 
| aws_cloudformation_stacks_selefra_id | string | X | X | fk to aws_cloudformation_stacks.selefra_id | 
| region | string | X | √ |  | 
| resource_status | string | X | √ |  | 
| drift_information | json | X | √ |  | 
| module_info | json | X | √ |  | 
| resource_status_reason | string | X | √ |  | 
| selefra_id | string | √ | √ | random id | 


