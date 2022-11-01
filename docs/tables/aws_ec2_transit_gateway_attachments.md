# Table: aws_ec2_transit_gateway_attachments

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| region | string | X | √ |  | 
| transit_gateway_arn | string | X | √ |  | 
| resource_type | string | X | √ |  | 
| creation_time | timestamp | X | √ |  | 
| resource_id | string | X | √ |  | 
| resource_owner_id | string | X | √ |  | 
| transit_gateway_owner_id | string | X | √ |  | 
| aws_ec2_transit_gateways_selefra_id | string | X | X | fk to aws_ec2_transit_gateways.selefra_id | 
| tags | json | X | √ |  | 
| association | json | X | √ |  | 
| transit_gateway_attachment_id | string | X | √ |  | 
| account_id | string | X | √ |  | 
| state | string | X | √ |  | 
| transit_gateway_id | string | X | √ |  | 
| selefra_id | string | √ | √ | random id | 


