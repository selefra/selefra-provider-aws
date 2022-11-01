# Table: aws_wafv2_ipsets

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| arn | string | √ | √ |  | 
| name | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| tags | json | X | √ |  | 
| description | string | X | √ |  | 
| addresses | ip_array | X | √ |  | 
| ip_address_version | string | X | √ |  | 
| id | string | X | √ |  | 


