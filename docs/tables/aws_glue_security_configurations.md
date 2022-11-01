# Table: aws_glue_security_configurations

## Primary Keys 

```
account_id, region, name
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| name | string | X | √ |  | 
| created_time_stamp | timestamp | X | √ |  | 
| encryption_configuration | json | X | √ |  | 


