# Table: aws_lightsail_alarms

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| comparison_operator | string | X | √ |  | 
| evaluation_periods | int | X | √ |  | 
| period | int | X | √ |  | 
| resource_type | string | X | √ |  | 
| state | string | X | √ |  | 
| unit | string | X | √ |  | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| name | string | X | √ |  | 
| notification_enabled | bool | X | √ |  | 
| treat_missing_data | string | X | √ |  | 
| monitored_resource_info | json | X | √ |  | 
| notification_triggers | string_array | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| metric_name | string | X | √ |  | 
| statistic | string | X | √ |  | 
| support_code | string | X | √ |  | 
| arn | string | √ | √ |  | 
| contact_protocols | string_array | X | √ |  | 
| created_at | timestamp | X | √ |  | 
| datapoints_to_alarm | int | X | √ |  | 
| location | json | X | √ |  | 
| threshold | float | X | √ |  | 


