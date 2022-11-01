# Table: aws_lightsail_disks

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| gb_in_use | int | X | √ |  | 
| is_system_disk | bool | X | √ |  | 
| path | string | X | √ |  | 
| state | string | X | √ |  | 
| arn | string | √ | √ |  | 
| tags | json | X | √ |  | 
| created_at | timestamp | X | √ |  | 
| size_in_gb | int | X | √ |  | 
| account_id | string | X | √ |  | 
| iops | int | X | √ |  | 
| resource_type | string | X | √ |  | 
| attachment_state | string | X | √ |  | 
| is_attached | bool | X | √ |  | 
| region | string | X | √ |  | 
| add_ons | json | X | √ |  | 
| attached_to | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| location | json | X | √ |  | 
| name | string | X | √ |  | 
| support_code | string | X | √ |  | 


