# Table: aws_iot_things

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| region | string | X | √ |  | 
| principals | string_array | X | √ |  | 
| arn | string | √ | √ |  | 
| thing_type_name | string | X | √ |  | 
| account_id | string | X | √ |  | 
| attributes | json | X | √ |  | 
| thing_name | string | X | √ |  | 
| version | int | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 


