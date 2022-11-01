# Table: aws_iot_billing_groups

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| billing_group_id | string | X | √ |  | 
| billing_group_name | string | X | √ |  | 
| billing_group_properties | json | X | √ |  | 
| arn | string | √ | √ |  | 
| region | string | X | √ |  | 
| things_in_group | string_array | X | √ |  | 
| tags | json | X | √ |  | 
| billing_group_metadata | json | X | √ |  | 
| version | int | X | √ |  | 
| result_metadata | json | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| account_id | string | X | √ |  | 


