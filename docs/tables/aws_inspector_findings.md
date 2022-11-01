# Table: aws_inspector_findings

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| attributes | json | X | √ |  | 
| service | string | X | √ |  | 
| account_id | string | X | √ |  | 
| description | string | X | √ |  | 
| schema_version | int | X | √ |  | 
| title | string | X | √ |  | 
| arn | string | √ | √ |  | 
| user_attributes | json | X | √ |  | 
| created_at | timestamp | X | √ |  | 
| asset_attributes | json | X | √ |  | 
| id | string | X | √ |  | 
| indicator_of_compromise | bool | X | √ |  | 
| numeric_severity | float | X | √ |  | 
| service_attributes | json | X | √ |  | 
| severity | string | X | √ |  | 
| region | string | X | √ |  | 
| updated_at | timestamp | X | √ |  | 
| asset_type | string | X | √ |  | 
| confidence | int | X | √ |  | 
| recommendation | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 


