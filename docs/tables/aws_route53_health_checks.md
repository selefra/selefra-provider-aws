# Table: aws_route53_health_checks

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| arn | string | √ | √ |  | 
| tags | json | X | √ | `The tags associated with the health check.` | 
| cloud_watch_alarm_configuration_dimensions | json | X | √ |  | 
| health_check_version | int | X | √ |  | 
| cloud_watch_alarm_configuration | json | X | √ |  | 
| linked_service | json | X | √ |  | 
| account_id | string | X | √ |  | 
| caller_reference | string | X | √ |  | 
| health_check_config | json | X | √ |  | 
| id | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 


