# Table: aws_glacier_vaults

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| region | string | X | √ |  | 
| last_inventory_date | string | X | √ |  | 
| number_of_archives | int | X | √ |  | 
| vault_name | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| account_id | string | X | √ |  | 
| tags | json | X | √ |  | 
| arn | string | √ | √ |  | 
| creation_date | string | X | √ |  | 
| size_in_bytes | int | X | √ |  | 


