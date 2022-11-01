# Table: aws_lightsail_static_ips

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| account_id | string | X | √ |  | 
| is_attached | bool | X | √ |  | 
| support_code | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| region | string | X | √ |  | 
| arn | string | √ | √ |  | 
| attached_to | string | X | √ |  | 
| created_at | timestamp | X | √ |  | 
| ip_address | string | X | √ |  | 
| location | json | X | √ |  | 
| name | string | X | √ |  | 
| resource_type | string | X | √ |  | 


