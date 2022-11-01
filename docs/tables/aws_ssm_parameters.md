# Table: aws_ssm_parameters

## Primary Keys 

```
account_id, region, name
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| policies | json | X | √ |  | 
| version | int | X | √ |  | 
| last_modified_user | string | X | √ |  | 
| last_modified_date | timestamp | X | √ |  | 
| type | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| description | string | X | √ |  | 
| name | string | X | √ | `The parameter name` | 
| account_id | string | X | √ | `The AWS Account ID of the resource` | 
| allowed_pattern | string | X | √ |  | 
| data_type | string | X | √ |  | 
| key_id | string | X | √ |  | 
| tier | string | X | √ |  | 
| region | string | X | √ | `The AWS Region of the resource` | 


