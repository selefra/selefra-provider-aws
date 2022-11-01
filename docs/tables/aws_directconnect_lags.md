# Table: aws_directconnect_lags

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| aws_device_v2 | string | X | √ |  | 
| encryption_mode | string | X | √ |  | 
| lag_state | string | X | √ |  | 
| mac_sec_capable | bool | X | √ |  | 
| region | string | X | √ |  | 
| tags | json | X | √ |  | 
| allows_hosted_connections | bool | X | √ |  | 
| aws_logical_device_id | string | X | √ |  | 
| location | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| account_id | string | X | √ |  | 
| aws_device | string | X | √ |  | 
| connections | json | X | √ |  | 
| has_logical_redundancy | string | X | √ |  | 
| lag_name | string | X | √ |  | 
| minimum_links | int | X | √ |  | 
| owner_account | string | X | √ |  | 
| arn | string | √ | √ |  | 
| id | string | X | √ |  | 
| connections_bandwidth | string | X | √ |  | 
| jumbo_frame_capable | bool | X | √ |  | 
| mac_sec_keys | json | X | √ |  | 
| number_of_connections | int | X | √ |  | 
| provider_name | string | X | √ |  | 


