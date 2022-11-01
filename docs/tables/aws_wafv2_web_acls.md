# Table: aws_wafv2_web_acls

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| resources_for_web_acl | string_array | X | √ |  | 
| custom_response_bodies | json | X | √ |  | 
| managed_by_firewall_manager | bool | X | √ |  | 
| post_process_firewall_manager_rule_groups | json | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| id | string | X | √ |  | 
| visibility_config | json | X | √ |  | 
| description | string | X | √ |  | 
| tags | json | X | √ |  | 
| pre_process_firewall_manager_rule_groups | json | X | √ |  | 
| rules | json | X | √ |  | 
| logging_configuration | json | X | √ |  | 
| label_namespace | string | X | √ |  | 
| arn | string | √ | √ |  | 
| default_action | json | X | √ |  | 
| name | string | X | √ |  | 
| capacity | int | X | √ |  | 
| captcha_config | json | X | √ |  | 


