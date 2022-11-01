# Table: aws_glue_database_tables

## Primary Keys 

```
database_arn, name
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| name | string | X | √ |  | 
| catalog_id | string | X | √ |  | 
| description | string | X | √ |  | 
| partition_keys | json | X | √ |  | 
| table_type | string | X | √ |  | 
| version_id | string | X | √ |  | 
| created_by | string | X | √ |  | 
| owner | string | X | √ |  | 
| parameters | json | X | √ |  | 
| retention | int | X | √ |  | 
| storage_descriptor | json | X | √ |  | 
| aws_glue_databases_selefra_id | string | X | X | fk to aws_glue_databases.selefra_id | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| is_registered_with_lake_formation | bool | X | √ |  | 
| view_expanded_text | string | X | √ |  | 
| view_original_text | string | X | √ |  | 
| database_arn | string | X | √ |  | 
| create_time | timestamp | X | √ |  | 
| database_name | string | X | √ |  | 
| last_access_time | timestamp | X | √ |  | 
| last_analyzed_time | timestamp | X | √ |  | 
| target_table | json | X | √ |  | 
| update_time | timestamp | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 


