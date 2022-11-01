# Table: aws_athena_work_groups

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| configuration | json | X | √ |  | 
| description | string | X | √ |  | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| arn | string | √ | √ |  | 
| tags | json | X | √ |  | 
| name | string | X | √ |  | 
| creation_time | timestamp | X | √ |  | 
| state | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 


