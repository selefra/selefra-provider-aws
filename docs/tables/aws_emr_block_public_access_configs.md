# Table: aws_emr_block_public_access_configs

## Primary Keys 

```
account_id, region
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| result_metadata | json | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| block_public_access_configuration | json | X | √ |  | 
| block_public_access_configuration_metadata | json | X | √ |  | 


