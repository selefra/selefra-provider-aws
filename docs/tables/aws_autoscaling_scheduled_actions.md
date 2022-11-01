# Table: aws_autoscaling_scheduled_actions

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| scheduled_action_name | string | X | √ |  | 
| arn | string | √ | √ |  | 
| auto_scaling_group_name | string | X | √ |  | 
| desired_capacity | int | X | √ |  | 
| end_time | timestamp | X | √ |  | 
| recurrence | string | X | √ |  | 
| start_time | timestamp | X | √ |  | 
| time | timestamp | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| region | string | X | √ |  | 
| min_size | int | X | √ |  | 
| account_id | string | X | √ |  | 
| time_zone | string | X | √ |  | 
| max_size | int | X | √ |  | 


