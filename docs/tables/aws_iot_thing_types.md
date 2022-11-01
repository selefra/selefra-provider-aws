# Table: aws_iot_thing_types

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| arn | string | √ | √ |  | 
| thing_type_metadata | json | X | √ |  | 
| thing_type_name | string | X | √ |  | 
| thing_type_properties | json | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| tags | json | X | √ |  | 


