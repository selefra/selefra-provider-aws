# Table: aws_ec2_flow_logs

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| creation_time | timestamp | X | √ |  | 
| flow_log_id | string | X | √ |  | 
| flow_log_status | string | X | √ |  | 
| traffic_type | string | X | √ |  | 
| tags | json | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| account_id | string | X | √ |  | 
| deliver_cross_account_role | string | X | √ |  | 
| deliver_logs_permission_arn | string | X | √ |  | 
| deliver_logs_status | string | X | √ |  | 
| log_destination | string | X | √ |  | 
| max_aggregation_interval | int | X | √ |  | 
| destination_options | json | X | √ |  | 
| log_destination_type | string | X | √ |  | 
| region | string | X | √ |  | 
| arn | string | √ | √ |  | 
| deliver_logs_error_message | string | X | √ |  | 
| log_format | string | X | √ |  | 
| log_group_name | string | X | √ |  | 
| resource_id | string | X | √ |  | 


