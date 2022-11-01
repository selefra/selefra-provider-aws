# Table: aws_wafregional_rate_based_rules

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| account_id | string | X | √ |  | 
| tags | json | X | √ |  | 
| match_predicates | json | X | √ |  | 
| rule_id | string | X | √ |  | 
| metric_name | string | X | √ |  | 
| name | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| region | string | X | √ |  | 
| arn | string | √ | √ |  | 
| rate_key | string | X | √ |  | 
| rate_limit | int | X | √ |  | 


