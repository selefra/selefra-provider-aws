# Table: aws_iot_ca_certificates

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| certificate_mode | string | X | √ |  | 
| creation_date | timestamp | X | √ |  | 
| generation_id | string | X | √ |  | 
| status | string | X | √ |  | 
| validity | json | X | √ |  | 
| arn | string | √ | √ |  | 
| certificate_id | string | X | √ |  | 
| auto_registration_status | string | X | √ |  | 
| certificate_pem | string | X | √ |  | 
| customer_version | int | X | √ |  | 
| owned_by | string | X | √ |  | 
| account_id | string | X | √ |  | 
| certificates | string_array | X | √ |  | 
| region | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| last_modified_date | timestamp | X | √ |  | 


