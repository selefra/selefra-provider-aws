# Table: aws_glue_crawlers

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| creation_time | timestamp | X | √ |  | 
| last_updated | timestamp | X | √ |  | 
| lineage_configuration | json | X | √ |  | 
| role | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| region | string | X | √ |  | 
| crawl_elapsed_time | int | X | √ |  | 
| lake_formation_configuration | json | X | √ |  | 
| recrawl_policy | json | X | √ |  | 
| version | int | X | √ |  | 
| classifiers | string_array | X | √ |  | 
| crawler_security_configuration | string | X | √ |  | 
| configuration | string | X | √ |  | 
| description | string | X | √ |  | 
| state | string | X | √ |  | 
| table_prefix | string | X | √ |  | 
| account_id | string | X | √ |  | 
| arn | string | √ | √ |  | 
| last_crawl | json | X | √ |  | 
| name | string | X | √ |  | 
| schedule | json | X | √ |  | 
| schema_change_policy | json | X | √ |  | 
| targets | json | X | √ |  | 
| tags | json | X | √ |  | 
| database_name | string | X | √ |  | 


