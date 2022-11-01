# Table: aws_route53_hosted_zones

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| id | string | X | √ |  | 
| config | json | X | √ |  | 
| linked_service | json | X | √ |  | 
| tags | json | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| arn | string | √ | √ |  | 
| caller_reference | string | X | √ |  | 
| resource_record_set_count | int | X | √ |  | 
| delegation_set_id | string | X | √ |  | 
| vpcs | json | X | √ |  | 
| account_id | string | X | √ |  | 
| name | string | X | √ |  | 


