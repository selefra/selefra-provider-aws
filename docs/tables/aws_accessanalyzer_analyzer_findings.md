# Table: aws_accessanalyzer_analyzer_findings

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| analyzed_at | timestamp | X | √ |  | 
| updated_at | timestamp | X | √ |  | 
| resource | string | X | √ |  | 
| arn | string | √ | √ |  | 
| resource_type | string | X | √ |  | 
| action | string_array | X | √ |  | 
| error | string | X | √ |  | 
| sources | json | X | √ |  | 
| aws_accessanalyzer_analyzers_selefra_id | string | X | X | fk to aws_accessanalyzer_analyzers.selefra_id | 
| condition | json | X | √ |  | 
| id | string | X | √ |  | 
| status | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| region | string | X | √ |  | 
| analyzer_arn | string | X | √ |  | 
| created_at | timestamp | X | √ |  | 
| resource_owner_account | string | X | √ |  | 
| is_public | bool | X | √ |  | 
| principal | json | X | √ |  | 
| account_id | string | X | √ |  | 


