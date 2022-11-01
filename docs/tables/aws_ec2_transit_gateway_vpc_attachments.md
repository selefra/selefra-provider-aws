# Table: aws_ec2_transit_gateway_vpc_attachments

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| account_id | string | X | √ |  | 
| subnet_ids | string_array | X | √ |  | 
| selefra_id | string | √ | √ | random id | 
| transit_gateway_id | string | X | √ |  | 
| vpc_owner_id | string | X | √ |  | 
| aws_ec2_transit_gateways_selefra_id | string | X | X | fk to aws_ec2_transit_gateways.selefra_id | 
| creation_time | timestamp | X | √ |  | 
| state | string | X | √ |  | 
| tags | json | X | √ |  | 
| options | json | X | √ |  | 
| transit_gateway_attachment_id | string | X | √ |  | 
| vpc_id | string | X | √ |  | 
| region | string | X | √ |  | 
| transit_gateway_arn | string | X | √ |  | 


