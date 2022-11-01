# Table: aws_cognito_user_pools

## Primary Keys 

```
account_id, region, id
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| account_recovery_setting | json | X | √ |  | 
| custom_domain | string | X | √ |  | 
| device_configuration | json | X | √ |  | 
| sms_configuration_failure | string | X | √ |  | 
| id | string | X | √ |  | 
| estimated_number_of_users | int | X | √ |  | 
| policies | json | X | √ |  | 
| sms_configuration | json | X | √ |  | 
| user_pool_add_ons | json | X | √ |  | 
| email_configuration_failure | string | X | √ |  | 
| email_verification_subject | string | X | √ |  | 
| mfa_configuration | string | X | √ |  | 
| schema_attributes | json | X | √ |  | 
| verification_message_template | json | X | √ |  | 
| last_modified_date | timestamp | X | √ |  | 
| sms_verification_message | string | X | √ |  | 
| status | string | X | √ |  | 
| username_configuration | json | X | √ |  | 
| user_pool_tags | json | X | √ |  | 
| username_attributes | string_array | X | √ |  | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| email_verification_message | string | X | √ |  | 
| sms_authentication_message | string | X | √ |  | 
| alias_attributes | string_array | X | √ |  | 
| creation_date | timestamp | X | √ |  | 
| lambda_config | json | X | √ |  | 
| auto_verified_attributes | string_array | X | √ |  | 
| name | string | X | √ |  | 
| user_attribute_update_settings | json | X | √ |  | 
| admin_create_user_config | json | X | √ |  | 
| arn | string | X | √ |  | 
| domain | string | X | √ |  | 
| email_configuration | json | X | √ |  | 


