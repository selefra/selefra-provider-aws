# Table: aws_fsx_data_repository_tasks

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| account_id | string | X | √ |  | 
| capacity_to_release | int | X | √ |  | 
| file_cache_id | string | X | √ |  | 
| arn | string | √ | √ |  | 
| lifecycle | string | X | √ |  | 
| failure_details | json | X | √ |  | 
| file_system_id | string | X | √ |  | 
| start_time | timestamp | X | √ |  | 
| tags | json | X | √ |  | 
| creation_time | timestamp | X | √ |  | 
| task_id | string | X | √ |  | 
| end_time | timestamp | X | √ |  | 
| status | json | X | √ |  | 
| region | string | X | √ |  | 
| type | string | X | √ |  | 
| paths | string_array | X | √ |  | 
| report | json | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 


