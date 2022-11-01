# Table: aws_lambda_function_aliases

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| account_id | string | X | √ |  | 
| function_version | string | X | √ |  | 
| routing_config | json | X | √ |  | 
| url_config | json | X | √ |  | 
| region | string | X | √ |  | 
| function_arn | string | X | √ |  | 
| arn | string | √ | √ |  | 
| description | string | X | √ |  | 
| name | string | X | √ |  | 
| revision_id | string | X | √ |  | 
| aws_lambda_functions_selefra_id | string | X | X | fk to aws_lambda_functions.selefra_id | 
| selefra_id | string | √ | √ | primary keys value md5 | 


