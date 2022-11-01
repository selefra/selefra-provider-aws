# Table: aws_glue_triggers

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| schedule | string | X | √ |  | 
| type | string | X | √ |  | 
| actions | json | X | √ |  | 
| description | string | X | √ |  | 
| id | string | X | √ |  | 
| name | string | X | √ |  | 
| region | string | X | √ |  | 
| tags | json | X | √ |  | 
| event_batching_condition | json | X | √ |  | 
| state | string | X | √ |  | 
| predicate | json | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| account_id | string | X | √ |  | 
| arn | string | √ | √ |  | 
| workflow_name | string | X | √ |  | 


