# Table: aws_glue_dev_endpoints

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| region | string | X | √ |  | 
| glue_version | string | X | √ |  | 
| number_of_workers | int | X | √ |  | 
| role_arn | string | X | √ |  | 
| created_timestamp | timestamp | X | √ |  | 
| last_modified_timestamp | timestamp | X | √ |  | 
| security_group_ids | string_array | X | √ |  | 
| yarn_endpoint_address | string | X | √ |  | 
| arn | string | √ | √ |  | 
| arguments | json | X | √ |  | 
| failure_reason | string | X | √ |  | 
| number_of_nodes | int | X | √ |  | 
| status | string | X | √ |  | 
| availability_zone | string | X | √ |  | 
| public_key | string | X | √ |  | 
| vpc_id | string | X | √ |  | 
| tags | json | X | √ |  | 
| private_address | string | X | √ |  | 
| public_keys | string_array | X | √ |  | 
| subnet_id | string | X | √ |  | 
| zeppelin_remote_spark_interpreter_port | int | X | √ |  | 
| account_id | string | X | √ |  | 
| last_update_status | string | X | √ |  | 
| security_configuration | string | X | √ |  | 
| worker_type | string | X | √ |  | 
| endpoint_name | string | X | √ |  | 
| extra_jars_s3_path | string | X | √ |  | 
| extra_python_libs_s3_path | string | X | √ |  | 
| public_address | string | X | √ |  | 


