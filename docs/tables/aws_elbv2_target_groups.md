# Table: aws_elbv2_target_groups

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| health_check_protocol | string | X | √ |  | 
| healthy_threshold_count | int | X | √ |  | 
| ip_address_type | string | X | √ |  | 
| load_balancer_arns | string_array | X | √ |  | 
| matcher | json | X | √ |  | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| arn | string | √ | √ |  | 
| protocol | string | X | √ |  | 
| target_group_name | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| health_check_port | string | X | √ |  | 
| tags | json | X | √ |  | 
| health_check_path | string | X | √ |  | 
| port | int | X | √ |  | 
| protocol_version | string | X | √ |  | 
| target_type | string | X | √ |  | 
| unhealthy_threshold_count | int | X | √ |  | 
| vpc_id | string | X | √ |  | 
| health_check_enabled | bool | X | √ |  | 
| health_check_interval_seconds | int | X | √ |  | 
| health_check_timeout_seconds | int | X | √ |  | 


