# Table: aws_shield_protections

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| health_check_ids | string_array | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| arn | string | √ | √ |  | 
| tags | json | X | √ |  | 
| id | string | X | √ |  | 
| name | string | X | √ |  | 
| resource_arn | string | X | √ |  | 
| account_id | string | X | √ |  | 
| application_layer_automatic_response_configuration | json | X | √ |  | 


