# Table: aws_iam_accounts

## Primary Keys 

```
account_id
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| groups_per_user_quota | int | X | √ |  | 
| mfa_devices | int | X | √ |  | 
| group_policy_size_quota | int | X | √ |  | 
| users | int | X | √ |  | 
| server_certificates | int | X | √ |  | 
| account_id | string | √ | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| account_signing_certificates_present | bool | X | √ |  | 
| versions_per_policy_quota | int | X | √ |  | 
| global_endpoint_token_version | int | X | √ |  | 
| user_policy_size_quota | int | X | √ |  | 
| mfa_devices_in_use | int | X | √ |  | 
| attached_policies_per_group_quota | int | X | √ |  | 
| policy_size_quota | int | X | √ |  | 
| policy_versions_in_use | int | X | √ |  | 
| aliases | string_array | X | √ |  | 
| access_keys_per_user_quota | int | X | √ |  | 
| groups_quota | int | X | √ |  | 
| server_certificates_quota | int | X | √ |  | 
| account_access_keys_present | bool | X | √ |  | 
| attached_policies_per_user_quota | int | X | √ |  | 
| policies_quota | int | X | √ |  | 
| users_quota | int | X | √ |  | 
| account_mfa_enabled | bool | X | √ |  | 
| attached_policies_per_role_quota | int | X | √ |  | 
| policies | int | X | √ |  | 
| policy_versions_in_use_quota | int | X | √ |  | 
| signing_certificates_per_user_quota | int | X | √ |  | 
| groups | int | X | √ |  | 


