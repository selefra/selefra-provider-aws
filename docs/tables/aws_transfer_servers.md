# Table: aws_transfer_servers

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| region | string | X | √ |  | 
| endpoint_type | string | X | √ |  | 
| host_key_fingerprint | string | X | √ |  | 
| logging_role | string | X | √ |  | 
| server_id | string | X | √ |  | 
| protocols | string_array | X | √ |  | 
| security_policy_name | string | X | √ |  | 
| account_id | string | X | √ |  | 
| arn | string | √ | √ |  | 
| identity_provider_details | json | X | √ |  | 
| post_authentication_login_banner | string | X | √ |  | 
| pre_authentication_login_banner | string | X | √ |  | 
| protocol_details | json | X | √ |  | 
| user_count | int | X | √ |  | 
| workflow_details | json | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| certificate | string | X | √ |  | 
| identity_provider_type | string | X | √ |  | 
| state | string | X | √ |  | 
| tags | json | X | √ | `Specifies the key-value pairs that you can use to search for and group servers that were assigned to the server that was described` | 
| domain | string | X | √ |  | 
| endpoint_details | json | X | √ |  | 


