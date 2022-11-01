# Table: aws_ec2_transit_gateway_route_tables

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| default_association_route_table | bool | X | √ |  | 
| state | string | X | √ |  | 
| transit_gateway_route_table_id | string | X | √ |  | 
| aws_ec2_transit_gateways_selefra_id | string | X | X | fk to aws_ec2_transit_gateways.selefra_id | 
| region | string | X | √ |  | 
| transit_gateway_arn | string | X | √ |  | 
| creation_time | timestamp | X | √ |  | 
| default_propagation_route_table | bool | X | √ |  | 
| transit_gateway_id | string | X | √ |  | 
| selefra_id | string | √ | √ | random id | 
| account_id | string | X | √ |  | 
| tags | json | X | √ |  | 


