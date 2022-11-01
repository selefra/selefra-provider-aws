# Table: aws_qldb_ledgers

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| encryption_description | json | X | √ |  | 
| name | string | X | √ |  | 
| permissions_mode | string | X | √ |  | 
| state | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| tags | json | X | √ | `The tags associated with the pipeline.` | 
| arn | string | √ | √ |  | 
| creation_date_time | timestamp | X | √ |  | 
| deletion_protection | bool | X | √ |  | 
| result_metadata | json | X | √ |  | 


