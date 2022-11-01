# Table: aws_iam_virtual_mfa_devices

## Primary Keys 

```
serial_number
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| base32_string_seed | int_array | X | √ |  | 
| qr_code_png | int_array | X | √ |  | 
| serial_number | string | √ | √ |  | 
| tags | json | X | √ |  | 
| enable_date | timestamp | X | √ |  | 
| user | json | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| account_id | string | X | √ |  | 
| user_tags | json | X | √ |  | 


