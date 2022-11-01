# Table: aws_directconnect_connections

## Primary Keys 

```
arn, id
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| connection_name | string | X | √ |  | 
| encryption_mode | string | X | √ |  | 
| jumbo_frame_capable | bool | X | √ |  | 
| owner_account | string | X | √ |  | 
| aws_logical_device_id | string | X | √ |  | 
| tags | json | X | √ |  | 
| has_logical_redundancy | string | X | √ |  | 
| lag_id | string | X | √ |  | 
| mac_sec_keys | json | X | √ |  | 
| port_encryption_status | string | X | √ |  | 
| vlan | int | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| id | string | X | √ |  | 
| region | string | X | √ |  | 
| arn | string | X | √ |  | 
| aws_device_v2 | string | X | √ |  | 
| bandwidth | string | X | √ |  | 
| location | string | X | √ |  | 
| partner_name | string | X | √ |  | 
| provider_name | string | X | √ |  | 
| account_id | string | X | √ |  | 
| connection_state | string | X | √ |  | 
| loa_issue_time | timestamp | X | √ |  | 
| mac_sec_capable | bool | X | √ |  | 
| aws_device | string | X | √ |  | 


