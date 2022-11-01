# Table: aws_lightsail_instance_snapshots

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| account_id | string | X | √ |  | 
| from_attached_disks | json | X | √ |  | 
| from_blueprint_id | string | X | √ |  | 
| from_instance_name | string | X | √ |  | 
| support_code | string | X | √ |  | 
| region | string | X | √ |  | 
| arn | string | √ | √ |  | 
| created_at | timestamp | X | √ |  | 
| progress | string | X | √ |  | 
| state | string | X | √ |  | 
| tags | json | X | √ |  | 
| name | string | X | √ |  | 
| size_in_gb | int | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| from_bundle_id | string | X | √ |  | 
| from_instance_arn | string | X | √ |  | 
| is_from_auto_snapshot | bool | X | √ |  | 
| location | json | X | √ |  | 
| resource_type | string | X | √ |  | 


