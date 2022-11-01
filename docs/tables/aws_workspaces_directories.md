# Table: aws_workspaces_directories

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| workspace_access_properties | json | X | √ |  | 
| workspace_creation_properties | json | X | √ |  | 
| directory_id | string | X | √ |  | 
| directory_name | string | X | √ |  | 
| directory_type | string | X | √ |  | 
| ip_group_ids | string_array | X | √ |  | 
| state | string | X | √ |  | 
| arn | string | √ | √ |  | 
| dns_ip_addresses | string_array | X | √ |  | 
| iam_role_id | string | X | √ |  | 
| workspace_security_group_id | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| account_id | string | X | √ |  | 
| alias | string | X | √ |  | 
| customer_user_name | string | X | √ |  | 
| saml_properties | json | X | √ |  | 
| selfservice_permissions | json | X | √ |  | 
| registration_code | string | X | √ |  | 
| subnet_ids | string_array | X | √ |  | 
| tenancy | string | X | √ |  | 


