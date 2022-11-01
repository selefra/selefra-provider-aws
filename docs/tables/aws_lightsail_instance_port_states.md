# Table: aws_lightsail_instance_port_states

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| ipv6_cidrs | string_array | X | √ |  | 
| protocol | string | X | √ |  | 
| state | string | X | √ |  | 
| to_port | int | X | √ |  | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| cidrs | string_array | X | √ |  | 
| from_port | int | X | √ |  | 
| selefra_id | string | √ | √ | random id | 
| aws_lightsail_instances_selefra_id | string | X | X | fk to aws_lightsail_instances.selefra_id | 
| instance_arn | string | X | √ |  | 
| cidr_list_aliases | string_array | X | √ |  | 


