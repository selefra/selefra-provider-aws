# Table: aws_fsx_volumes

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| ontap_configuration | json | X | √ |  | 
| open_zfs_configuration | json | X | √ |  | 
| volume_type | string | X | √ |  | 
| region | string | X | √ |  | 
| arn | string | √ | √ |  | 
| tags | json | X | √ |  | 
| administrative_actions | json | X | √ |  | 
| file_system_id | string | X | √ |  | 
| name | string | X | √ |  | 
| volume_id | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| account_id | string | X | √ |  | 
| creation_time | timestamp | X | √ |  | 
| lifecycle | string | X | √ |  | 
| lifecycle_transition_reason | json | X | √ |  | 


