# Table: aws_applicationautoscaling_policies

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| policy_name | string | X | √ |  | 
| policy_type | string | X | √ |  | 
| scalable_dimension | string | X | √ |  | 
| service_namespace | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| account_id | string | X | √ |  | 
| arn | string | √ | √ |  | 
| creation_time | timestamp | X | √ |  | 
| resource_id | string | X | √ |  | 
| alarms | json | X | √ |  | 
| step_scaling_policy_configuration | json | X | √ |  | 
| target_tracking_scaling_policy_configuration | json | X | √ |  | 
| region | string | X | √ |  | 


