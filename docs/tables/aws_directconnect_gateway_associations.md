# Table: aws_directconnect_gateway_associations

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| account_id | string | X | √ |  | 
| gateway_arn | string | X | √ |  | 
| gateway_id | string | X | √ |  | 
| association_id | string | X | √ |  | 
| direct_connect_gateway_id | string | X | √ |  | 
| virtual_gateway_id | string | X | √ |  | 
| region | string | X | √ |  | 
| aws_directconnect_gateways_selefra_id | string | X | X | fk to aws_directconnect_gateways.selefra_id | 
| allowed_prefixes_to_direct_connect_gateway | json | X | √ |  | 
| associated_gateway | json | X | √ |  | 
| association_state | string | X | √ |  | 
| direct_connect_gateway_owner_account | string | X | √ |  | 
| state_change_error | string | X | √ |  | 
| virtual_gateway_owner_account | string | X | √ |  | 
| virtual_gateway_region | string | X | √ |  | 
| selefra_id | string | √ | √ | random id | 


