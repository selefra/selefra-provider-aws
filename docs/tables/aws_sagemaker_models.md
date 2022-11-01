# Table: aws_sagemaker_models

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| region | string | X | √ |  | 
| containers | json | X | √ |  | 
| tags | json | X | √ | `The tags associated with the model.` | 
| execution_role_arn | string | X | √ |  | 
| model_name | string | X | √ |  | 
| account_id | string | X | √ |  | 
| creation_time | timestamp | X | √ |  | 
| vpc_config | json | X | √ |  | 
| result_metadata | json | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| arn | string | √ | √ |  | 
| enable_network_isolation | bool | X | √ |  | 
| inference_execution_config | json | X | √ |  | 
| primary_container | json | X | √ |  | 


