# Table: aws_wafregional_rules

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| predicates | json | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| arn | string | √ | √ |  | 
| tags | json | X | √ | `Rule tags.` | 
| rule_id | string | X | √ |  | 
| metric_name | string | X | √ |  | 
| name | string | X | √ |  | 


