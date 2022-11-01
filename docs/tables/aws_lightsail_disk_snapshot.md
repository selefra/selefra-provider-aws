# Table: aws_lightsail_disk_snapshot

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| is_from_auto_snapshot | bool | X | √ |  | 
| resource_type | string | X | √ |  | 
| support_code | string | X | √ |  | 
| created_at | timestamp | X | √ |  | 
| from_instance_name | string | X | √ |  | 
| progress | string | X | √ |  | 
| state | string | X | √ |  | 
| from_disk_arn | string | X | √ |  | 
| from_disk_name | string | X | √ |  | 
| from_instance_arn | string | X | √ |  | 
| location | json | X | √ |  | 
| size_in_gb | int | X | √ |  | 
| aws_lightsail_disks_selefra_id | string | X | X | fk to aws_lightsail_disks.selefra_id | 
| tags | json | X | √ |  | 
| region | string | X | √ |  | 
| disk_arn | string | X | √ |  | 
| arn | string | X | √ |  | 
| name | string | X | √ |  | 
| selefra_id | string | √ | √ | random id | 
| account_id | string | X | √ |  | 


