# Table: aws_athena_data_catalog_database_tables

## Primary Keys 

```
data_catalog_arn, data_catalog_database_name, name
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| create_time | timestamp | X | √ |  | 
| last_access_time | timestamp | X | √ |  | 
| partition_keys | json | X | √ |  | 
| region | string | X | √ |  | 
| data_catalog_database_name | string | X | √ |  | 
| columns | json | X | √ |  | 
| parameters | json | X | √ |  | 
| table_type | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| aws_athena_data_catalog_databases_selefra_id | string | X | X | fk to aws_athena_data_catalog_databases.selefra_id | 
| account_id | string | X | √ |  | 
| data_catalog_arn | string | X | √ |  | 
| name | string | X | √ |  | 


