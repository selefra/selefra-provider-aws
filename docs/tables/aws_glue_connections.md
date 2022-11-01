# Table: aws_glue_connections

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| description | string | X | √ |  | 
| last_updated_by | string | X | √ |  | 
| last_updated_time | timestamp | X | √ |  | 
| match_criteria | string_array | X | √ |  | 
| connection_type | string | X | √ |  | 
| region | string | X | √ |  | 
| arn | string | √ | √ |  | 
| connection_properties | json | X | √ |  | 
| creation_time | timestamp | X | √ |  | 
| name | string | X | √ |  | 
| physical_connection_requirements | json | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| account_id | string | X | √ |  | 


