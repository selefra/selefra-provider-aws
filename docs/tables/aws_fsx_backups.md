# Table: aws_fsx_backups

## Primary Keys 

```
account_id, region, id
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| lifecycle | string | X | √ |  | 
| type | string | X | √ |  | 
| directory_information | json | X | √ |  | 
| kms_key_id | string | X | √ |  | 
| resource_arn | string | X | √ |  | 
| tags | json | X | √ |  | 
| source_backup_id | string | X | √ |  | 
| region | string | X | √ |  | 
| file_system | json | X | √ |  | 
| failure_details | json | X | √ |  | 
| owner_id | string | X | √ |  | 
| source_backup_region | string | X | √ |  | 
| volume | json | X | √ |  | 
| creation_time | timestamp | X | √ |  | 
| id | string | X | √ |  | 
| progress_percent | int | X | √ |  | 
| resource_type | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| account_id | string | X | √ |  | 


