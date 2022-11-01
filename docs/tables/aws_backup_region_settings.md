# Table: aws_backup_region_settings

## Primary Keys 

```
account_id, region
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| resource_type_management_preference | json | X | √ |  | 
| resource_type_opt_in_preference | json | X | √ |  | 
| result_metadata | json | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 


