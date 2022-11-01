# Table: aws_cloudformation_stacks

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| capabilities | string_array | X | √ |  | 
| deletion_time | timestamp | X | √ |  | 
| notification_ar_ns | string_array | X | √ |  | 
| outputs | json | X | √ |  | 
| root_id | string | X | √ |  | 
| stack_status_reason | string | X | √ |  | 
| parameters | json | X | √ |  | 
| parent_id | string | X | √ |  | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| tags | json | X | √ |  | 
| stack_status | string | X | √ |  | 
| drift_information | json | X | √ |  | 
| last_updated_time | timestamp | X | √ |  | 
| timeout_in_minutes | int | X | √ |  | 
| id | string | X | √ |  | 
| arn | string | √ | √ |  | 
| stack_name | string | X | √ |  | 
| disable_rollback | bool | X | √ |  | 
| enable_termination_protection | bool | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| creation_time | timestamp | X | √ |  | 
| change_set_id | string | X | √ |  | 
| description | string | X | √ |  | 
| role_arn | string | X | √ |  | 
| rollback_configuration | json | X | √ |  | 


