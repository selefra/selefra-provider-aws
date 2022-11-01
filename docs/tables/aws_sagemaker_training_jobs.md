# Table: aws_sagemaker_training_jobs

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| account_id | string | X | √ |  | 
| model_artifacts | json | X | √ |  | 
| debug_hook_config | json | X | √ |  | 
| final_metric_data_list | json | X | √ |  | 
| tuning_job_arn | string | X | √ |  | 
| warm_pool_status | json | X | √ |  | 
| result_metadata | json | X | √ |  | 
| arn | string | √ | √ |  | 
| secondary_status | string | X | √ |  | 
| training_job_name | string | X | √ |  | 
| billable_time_in_seconds | int | X | √ |  | 
| enable_network_isolation | bool | X | √ |  | 
| failure_reason | string | X | √ |  | 
| training_time_in_seconds | int | X | √ |  | 
| secondary_status_transitions | json | X | √ |  | 
| tags | json | X | √ | `The tags associated with the model.` | 
| resource_config | json | X | √ |  | 
| checkpoint_config | json | X | √ |  | 
| environment | json | X | √ |  | 
| hyper_parameters | json | X | √ |  | 
| last_modified_time | timestamp | X | √ |  | 
| profiler_config | json | X | √ |  | 
| algorithm_specification | json | X | √ |  | 
| experiment_config | json | X | √ |  | 
| profiler_rule_configurations | json | X | √ |  | 
| profiling_status | string | X | √ |  | 
| tensor_board_output_config | json | X | √ |  | 
| role_arn | string | X | √ |  | 
| creation_time | timestamp | X | √ |  | 
| training_job_status | string | X | √ |  | 
| debug_rule_configurations | json | X | √ |  | 
| input_data_config | json | X | √ |  | 
| labeling_job_arn | string | X | √ |  | 
| profiler_rule_evaluation_statuses | json | X | √ |  | 
| retry_strategy | json | X | √ |  | 
| training_end_time | timestamp | X | √ |  | 
| training_start_time | timestamp | X | √ |  | 
| vpc_config | json | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| region | string | X | √ |  | 
| stopping_condition | json | X | √ |  | 
| auto_ml_job_arn | string | X | √ |  | 
| debug_rule_evaluation_statuses | json | X | √ |  | 
| enable_inter_container_traffic_encryption | bool | X | √ |  | 
| enable_managed_spot_training | bool | X | √ |  | 
| output_data_config | json | X | √ |  | 


