# Table: aws_iot_thing_groups

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| region | string | X | √ |  | 
| query_string | string | X | √ |  | 
| thing_group_metadata | json | X | √ |  | 
| thing_group_name | string | X | √ |  | 
| version | int | X | √ |  | 
| status | string | X | √ |  | 
| result_metadata | json | X | √ |  | 
| policies | string_array | X | √ |  | 
| tags | json | X | √ |  | 
| index_name | string | X | √ |  | 
| thing_group_id | string | X | √ |  | 
| thing_group_properties | json | X | √ |  | 
| account_id | string | X | √ |  | 
| things_in_group | string_array | X | √ |  | 
| arn | string | √ | √ |  | 
| query_version | string | X | √ |  | 


