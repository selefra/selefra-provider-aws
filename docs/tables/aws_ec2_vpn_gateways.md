# Table: aws_ec2_vpn_gateways

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| type | string | X | √ |  | 
| vpc_attachments | json | X | √ |  | 
| region | string | X | √ |  | 
| arn | string | √ | √ |  | 
| amazon_side_asn | int | X | √ |  | 
| tags | json | X | √ |  | 
| vpn_gateway_id | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| account_id | string | X | √ |  | 
| availability_zone | string | X | √ |  | 
| state | string | X | √ |  | 


