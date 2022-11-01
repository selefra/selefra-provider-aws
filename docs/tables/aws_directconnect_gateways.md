# Table: aws_directconnect_gateways

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| region | string | X | √ |  | 
| arn | string | √ | √ |  | 
| id | string | X | √ |  | 
| direct_connect_gateway_state | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| account_id | string | X | √ |  | 
| amazon_side_asn | int | X | √ |  | 
| direct_connect_gateway_name | string | X | √ |  | 
| owner_account | string | X | √ |  | 
| state_change_error | string | X | √ |  | 


