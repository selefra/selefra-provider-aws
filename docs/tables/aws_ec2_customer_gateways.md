# Table: aws_ec2_customer_gateways

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| state | string | X | √ |  | 
| tags | json | X | √ |  | 
| type | string | X | √ |  | 
| region | string | X | √ |  | 
| arn | string | √ | √ |  | 
| bgp_asn | string | X | √ |  | 
| customer_gateway_id | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| account_id | string | X | √ |  | 
| certificate_arn | string | X | √ |  | 
| device_name | string | X | √ |  | 
| ip_address | string | X | √ |  | 


