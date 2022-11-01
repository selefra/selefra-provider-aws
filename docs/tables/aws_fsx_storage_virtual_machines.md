# Table: aws_fsx_storage_virtual_machines

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| creation_time | timestamp | X | √ |  | 
| file_system_id | string | X | √ |  | 
| lifecycle | string | X | √ |  | 
| root_volume_security_style | string | X | √ |  | 
| active_directory_configuration | json | X | √ |  | 
| uuid | string | X | √ |  | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| arn | string | √ | √ |  | 
| tags | json | X | √ |  | 
| subtype | string | X | √ |  | 
| endpoints | json | X | √ |  | 
| lifecycle_transition_reason | json | X | √ |  | 
| name | string | X | √ |  | 
| storage_virtual_machine_id | string | X | √ |  | 


