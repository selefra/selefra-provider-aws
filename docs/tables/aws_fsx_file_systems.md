# Table: aws_fsx_file_systems

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| kms_key_id | string | X | √ |  | 
| lifecycle | string | X | √ |  | 
| network_interface_ids | string_array | X | √ |  | 
| vpc_id | string | X | √ |  | 
| arn | string | √ | √ |  | 
| tags | json | X | √ |  | 
| administrative_actions | json | X | √ |  | 
| windows_configuration | json | X | √ |  | 
| account_id | string | X | √ |  | 
| creation_time | timestamp | X | √ |  | 
| dns_name | string | X | √ |  | 
| file_system_type_version | string | X | √ |  | 
| lustre_configuration | json | X | √ |  | 
| open_zfs_configuration | json | X | √ |  | 
| subnet_ids | string_array | X | √ |  | 
| region | string | X | √ |  | 
| failure_details | json | X | √ |  | 
| file_system_id | string | X | √ |  | 
| storage_capacity | int | X | √ |  | 
| storage_type | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| file_system_type | string | X | √ |  | 
| ontap_configuration | json | X | √ |  | 
| owner_id | string | X | √ |  | 


