# Table: aws_backup_plans

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| tags | json | X | √ |  | 
| advanced_backup_settings | json | X | √ |  | 
| deletion_date | timestamp | X | √ |  | 
| version_id | string | X | √ |  | 
| last_execution_date | timestamp | X | √ |  | 
| result_metadata | json | X | √ |  | 
| arn | string | √ | √ |  | 
| backup_plan | json | X | √ |  | 
| creation_date | timestamp | X | √ |  | 
| creator_request_id | string | X | √ |  | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| backup_plan_id | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 


