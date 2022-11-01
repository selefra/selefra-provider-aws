# Table: aws_autoscaling_group_scaling_policies

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| auto_scaling_group_name | string | X | √ |  | 
| estimated_instance_warmup | int | X | √ |  | 
| min_adjustment_step | int | X | √ |  | 
| policy_name | string | X | √ |  | 
| predictive_scaling_configuration | json | X | √ |  | 
| adjustment_type | string | X | √ |  | 
| region | string | X | √ |  | 
| arn | string | √ | √ |  | 
| alarms | json | X | √ |  | 
| metric_aggregation_type | string | X | √ |  | 
| scaling_adjustment | int | X | √ |  | 
| step_adjustments | json | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| account_id | string | X | √ |  | 
| enabled | bool | X | √ |  | 
| min_adjustment_magnitude | int | X | √ |  | 
| aws_autoscaling_groups_selefra_id | string | X | X | fk to aws_autoscaling_groups.selefra_id | 
| cooldown | int | X | √ |  | 
| policy_type | string | X | √ |  | 
| target_tracking_configuration | json | X | √ |  | 
| group_arn | string | X | √ |  | 


