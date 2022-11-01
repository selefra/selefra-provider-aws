# Table: aws_iot_jobs

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| job_process_details | json | X | √ |  | 
| last_updated_at | timestamp | X | √ |  | 
| namespace_id | string | X | √ |  | 
| reason_code | string | X | √ |  | 
| tags | json | X | √ |  | 
| comment | string | X | √ |  | 
| description | string | X | √ |  | 
| job_executions_rollout_config | json | X | √ |  | 
| job_template_arn | string | X | √ |  | 
| presigned_url_config | json | X | √ |  | 
| target_selection | string | X | √ |  | 
| created_at | timestamp | X | √ |  | 
| document_parameters | json | X | √ |  | 
| force_canceled | bool | X | √ |  | 
| is_concurrent | bool | X | √ |  | 
| arn | string | √ | √ |  | 
| status | string | X | √ |  | 
| timeout_config | json | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| job_executions_retry_config | json | X | √ |  | 
| job_id | string | X | √ |  | 
| targets | string_array | X | √ |  | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| abort_config | json | X | √ |  | 
| completed_at | timestamp | X | √ |  | 


