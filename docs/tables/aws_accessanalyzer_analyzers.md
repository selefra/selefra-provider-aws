# Table: aws_accessanalyzer_analyzers

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| arn | string | √ | √ |  | 
| created_at | timestamp | X | √ |  | 
| name | string | X | √ |  | 
| type | string | X | √ |  | 
| status_reason | json | X | √ |  | 
| tags | json | X | √ |  | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| last_resource_analyzed_at | timestamp | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| status | string | X | √ |  | 
| last_resource_analyzed | string | X | √ |  | 


