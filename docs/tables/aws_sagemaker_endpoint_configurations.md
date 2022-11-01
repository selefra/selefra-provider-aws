# Table: aws_sagemaker_endpoint_configurations

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| arn | string | √ | √ |  | 
| creation_time | timestamp | X | √ |  | 
| production_variants | json | X | √ |  | 
| data_capture_config | json | X | √ |  | 
| kms_key_id | string | X | √ |  | 
| result_metadata | json | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| tags | json | X | √ | `The tags associated with the model.` | 
| endpoint_config_name | string | X | √ |  | 
| async_inference_config | json | X | √ |  | 
| explainer_config | json | X | √ |  | 


