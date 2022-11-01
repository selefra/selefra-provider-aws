# Table: aws_glue_workflows

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| last_modified_on | timestamp | X | √ |  | 
| last_run | json | X | √ |  | 
| max_concurrent_runs | int | X | √ |  | 
| account_id | string | X | √ |  | 
| default_run_properties | json | X | √ |  | 
| description | string | X | √ |  | 
| blueprint_details | json | X | √ |  | 
| graph | json | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| region | string | X | √ |  | 
| arn | string | √ | √ |  | 
| tags | json | X | √ |  | 
| created_on | timestamp | X | √ |  | 
| name | string | X | √ |  | 


