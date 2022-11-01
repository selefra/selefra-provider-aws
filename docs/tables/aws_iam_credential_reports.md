# Table: aws_iam_credential_reports

## Primary Keys 

```
arn, user_creation_time
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| access_key_1_last_used_date | timestamp | X | √ |  | 
| password_last_used | timestamp | X | √ |  | 
| cert2_active | bool | X | √ |  | 
| access_key1_last_used_region | string | X | √ |  | 
| user_creation_time | timestamp | X | √ |  | 
| password_last_changed | timestamp | X | √ |  | 
| access_key_2_last_rotated | timestamp | X | √ |  | 
| cert_1_last_rotated | timestamp | X | √ |  | 
| access_key2_last_used_region | string | X | √ |  | 
| access_key2_last_used_service | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| arn | string | X | √ |  | 
| cert_2_last_rotated | timestamp | X | √ |  | 
| user | string | X | √ |  | 
| access_key1_active | bool | X | √ |  | 
| access_key2_active | bool | X | √ |  | 
| access_key1_last_used_service | string | X | √ |  | 
| password_next_rotation | timestamp | X | √ |  | 
| access_key_1_last_rotated | timestamp | X | √ |  | 
| access_key_2_last_used_date | timestamp | X | √ |  | 
| mfa_active | bool | X | √ |  | 
| password_status | string | X | √ |  | 
| cert1_active | bool | X | √ |  | 


