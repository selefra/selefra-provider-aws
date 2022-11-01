# Table: aws_sagemaker_notebook_instances

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| default_code_repository | string | X | √ |  | 
| last_modified_time | timestamp | X | √ |  | 
| notebook_instance_lifecycle_config_name | string | X | √ |  | 
| result_metadata | json | X | √ |  | 
| region | string | X | √ |  | 
| direct_internet_access | string | X | √ |  | 
| volume_size_in_gb | int | X | √ |  | 
| creation_time | timestamp | X | √ |  | 
| instance_type | string | X | √ |  | 
| tags | json | X | √ | `The tags associated with the notebook instance.` | 
| instance_metadata_service_configuration | json | X | √ |  | 
| role_arn | string | X | √ |  | 
| url | string | X | √ |  | 
| account_id | string | X | √ |  | 
| arn | string | √ | √ |  | 
| additional_code_repositories | string_array | X | √ |  | 
| platform_identifier | string | X | √ |  | 
| subnet_id | string | X | √ |  | 
| accelerator_types | string_array | X | √ |  | 
| kms_key_id | string | X | √ |  | 
| security_groups | string_array | X | √ |  | 
| root_access | string | X | √ |  | 
| failure_reason | string | X | √ |  | 
| network_interface_id | string | X | √ |  | 
| notebook_instance_name | string | X | √ |  | 
| notebook_instance_status | string | X | √ |  | 


