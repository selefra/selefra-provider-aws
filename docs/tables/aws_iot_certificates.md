# Table: aws_iot_certificates

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| certificate_id | string | X | √ |  | 
| certificate_mode | string | X | √ |  | 
| generation_id | string | X | √ |  | 
| last_modified_date | timestamp | X | √ |  | 
| arn | string | √ | √ |  | 
| ca_certificate_id | string | X | √ |  | 
| certificate_pem | string | X | √ |  | 
| previous_owned_by | string | X | √ |  | 
| account_id | string | X | √ |  | 
| policies | string_array | X | √ |  | 
| transfer_data | json | X | √ |  | 
| validity | json | X | √ |  | 
| creation_date | timestamp | X | √ |  | 
| status | string | X | √ |  | 
| owned_by | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| region | string | X | √ |  | 
| customer_version | int | X | √ |  | 


