# Table: aws_iot_security_profiles

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| security_profile_name | string | X | √ |  | 
| region | string | X | √ |  | 
| arn | string | √ | √ |  | 
| last_modified_date | timestamp | X | √ |  | 
| security_profile_description | string | X | √ |  | 
| result_metadata | json | X | √ |  | 
| account_id | string | X | √ |  | 
| tags | json | X | √ |  | 
| additional_metrics_to_retain | string_array | X | √ |  | 
| version | int | X | √ |  | 
| alert_targets | json | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| targets | string_array | X | √ |  | 
| additional_metrics_to_retain_v2 | json | X | √ |  | 
| behaviors | json | X | √ |  | 
| creation_date | timestamp | X | √ |  | 


