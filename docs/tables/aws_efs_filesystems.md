# Table: aws_efs_filesystems

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| name | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| account_id | string | X | √ |  | 
| life_cycle_state | string | X | √ |  | 
| size_in_bytes | json | X | √ |  | 
| creation_time | timestamp | X | √ |  | 
| owner_id | string | X | √ |  | 
| provisioned_throughput_in_mibps | float | X | √ |  | 
| throughput_mode | string | X | √ |  | 
| arn | string | √ | √ |  | 
| creation_token | string | X | √ |  | 
| kms_key_id | string | X | √ |  | 
| number_of_mount_targets | int | X | √ |  | 
| performance_mode | string | X | √ |  | 
| tags | json | X | √ |  | 
| availability_zone_id | string | X | √ |  | 
| availability_zone_name | string | X | √ |  | 
| region | string | X | √ |  | 
| backup_policy_status | string | X | √ |  | 
| file_system_id | string | X | √ |  | 
| encrypted | bool | X | √ |  | 


