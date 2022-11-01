# Table: aws_route53_delegation_sets

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| name_servers | string_array | X | √ |  | 
| caller_reference | string | X | √ |  | 
| id | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| account_id | string | X | √ |  | 
| arn | string | √ | √ | `The Amazon Resource Name (ARN) for the resource.` | 


