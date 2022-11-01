# Table: aws_workspaces_workspaces

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| subnet_id | string | X | √ |  | 
| account_id | string | X | √ |  | 
| error_code | string | X | √ |  | 
| ip_address | string | X | √ |  | 
| user_volume_encryption_enabled | bool | X | √ |  | 
| workspace_id | string | X | √ |  | 
| workspace_properties | json | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| arn | string | √ | √ |  | 
| root_volume_encryption_enabled | bool | X | √ |  | 
| user_name | string | X | √ |  | 
| volume_encryption_key | string | X | √ |  | 
| bundle_id | string | X | √ |  | 
| directory_id | string | X | √ |  | 
| state | string | X | √ |  | 
| computer_name | string | X | √ |  | 
| error_message | string | X | √ |  | 
| modification_states | json | X | √ |  | 


