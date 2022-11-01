# Table: aws_glue_jobs

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| tags | json | X | √ |  | 
| description | string | X | √ |  | 
| execution_property | json | X | √ |  | 
| max_capacity | float | X | √ |  | 
| name | string | X | √ |  | 
| worker_type | string | X | √ |  | 
| command | json | X | √ |  | 
| default_arguments | json | X | √ |  | 
| execution_class | string | X | √ |  | 
| security_configuration | string | X | √ |  | 
| account_id | string | X | √ |  | 
| arn | string | √ | √ |  | 
| role | string | X | √ |  | 
| created_on | timestamp | X | √ |  | 
| source_control_details | json | X | √ |  | 
| region | string | X | √ |  | 
| allocated_capacity | int | X | √ |  | 
| last_modified_on | timestamp | X | √ |  | 
| max_retries | int | X | √ |  | 
| notification_property | json | X | √ |  | 
| non_overridable_arguments | json | X | √ |  | 
| timeout | int | X | √ |  | 
| log_uri | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| code_gen_configuration_nodes | json | X | √ |  | 
| connections | json | X | √ |  | 
| glue_version | string | X | √ |  | 
| number_of_workers | int | X | √ |  | 


