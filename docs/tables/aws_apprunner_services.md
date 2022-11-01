# Table: aws_apprunner_services

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| observability_configuration | json | X | √ |  | 
| region | string | X | √ |  | 
| auto_scaling_configuration_summary | json | X | √ |  | 
| updated_at | timestamp | X | √ |  | 
| encryption_configuration | json | X | √ |  | 
| health_check_configuration | json | X | √ |  | 
| account_id | string | X | √ |  | 
| created_at | timestamp | X | √ |  | 
| source_configuration | json | X | √ |  | 
| status | string | X | √ |  | 
| service_id | string | X | √ |  | 
| service_url | string | X | √ |  | 
| network_configuration | json | X | √ |  | 
| service_name | string | X | √ |  | 
| deleted_at | timestamp | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| arn | string | √ | √ |  | 
| instance_configuration | json | X | √ |  | 


