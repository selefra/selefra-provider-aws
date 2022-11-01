# Table: aws_ssm_instances

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| last_ping_date_time | timestamp | X | √ |  | 
| last_successful_association_execution_date | timestamp | X | √ |  | 
| platform_type | string | X | √ |  | 
| source_type | string | X | √ |  | 
| arn | string | √ | √ |  | 
| activation_id | string | X | √ |  | 
| association_status | string | X | √ |  | 
| last_association_execution_date | timestamp | X | √ |  | 
| iam_role | string | X | √ |  | 
| instance_id | string | X | √ |  | 
| is_latest_version | bool | X | √ |  | 
| resource_type | string | X | √ |  | 
| registration_date | timestamp | X | √ |  | 
| source_id | string | X | √ |  | 
| association_overview | json | X | √ |  | 
| computer_name | string | X | √ |  | 
| ping_status | string | X | √ |  | 
| platform_name | string | X | √ |  | 
| name | string | X | √ |  | 
| platform_version | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| agent_version | string | X | √ |  | 
| ip_address | string | X | √ |  | 


