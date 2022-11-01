# Table: aws_glue_database_table_indexes

## Primary Keys 

```
database_arn, database_table_name, index_name
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| database_arn | string | X | √ |  | 
| database_table_name | string | X | √ |  | 
| index_name | string | X | √ |  | 
| index_status | string | X | √ |  | 
| backfill_errors | json | X | √ |  | 
| aws_glue_database_tables_selefra_id | string | X | X | fk to aws_glue_database_tables.selefra_id | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| account_id | string | X | √ |  | 
| keys | json | X | √ |  | 
| region | string | X | √ |  | 


