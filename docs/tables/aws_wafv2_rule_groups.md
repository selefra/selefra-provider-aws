# Table: aws_wafv2_rule_groups

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| tags | json | X | √ |  | 
| policy | json | X | √ |  | 
| available_labels | json | X | √ |  | 
| consumed_labels | json | X | √ |  | 
| label_namespace | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| arn | string | √ | √ |  | 
| id | string | X | √ |  | 
| custom_response_bodies | json | X | √ |  | 
| capacity | int | X | √ |  | 
| name | string | X | √ |  | 
| visibility_config | json | X | √ |  | 
| description | string | X | √ |  | 
| rules | json | X | √ |  | 


