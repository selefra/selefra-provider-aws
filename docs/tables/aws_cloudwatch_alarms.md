# Table: aws_cloudwatch_alarms

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| state_value | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| ok_actions | string_array | X | √ |  | 
| state_updated_timestamp | timestamp | X | √ |  | 
| unit | string | X | √ |  | 
| region | string | X | √ |  | 
| tags | json | X | √ |  | 
| actions_enabled | bool | X | √ |  | 
| alarm_actions | string_array | X | √ |  | 
| alarm_configuration_updated_timestamp | timestamp | X | √ |  | 
| metric_name | string | X | √ |  | 
| metrics | json | X | √ |  | 
| threshold_metric_id | string | X | √ |  | 
| account_id | string | X | √ |  | 
| dimensions | json | X | √ |  | 
| alarm_name | string | X | √ |  | 
| extended_statistic | string | X | √ |  | 
| insufficient_data_actions | string_array | X | √ |  | 
| alarm_description | string | X | √ |  | 
| threshold | float | X | √ |  | 
| arn | string | √ | √ |  | 
| datapoints_to_alarm | int | X | √ |  | 
| evaluate_low_sample_count_percentile | string | X | √ |  | 
| state_reason_data | string | X | √ |  | 
| treat_missing_data | string | X | √ |  | 
| statistic | string | X | √ |  | 
| comparison_operator | string | X | √ |  | 
| evaluation_periods | int | X | √ |  | 
| namespace | string | X | √ |  | 
| period | int | X | √ |  | 
| state_reason | string | X | √ |  | 


