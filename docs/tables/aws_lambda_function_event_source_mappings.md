# Table: aws_lambda_function_event_source_mappings

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| event_source_arn | string | X | √ |  | 
| parallelization_factor | int | X | √ |  | 
| queues | string_array | X | √ |  | 
| account_id | string | X | √ |  | 
| aws_lambda_functions_selefra_id | string | X | X | fk to aws_lambda_functions.selefra_id | 
| last_modified | timestamp | X | √ |  | 
| starting_position | string | X | √ |  | 
| starting_position_timestamp | timestamp | X | √ |  | 
| bisect_batch_on_function_error | bool | X | √ |  | 
| amazon_managed_kafka_event_source_config | json | X | √ |  | 
| batch_size | int | X | √ |  | 
| source_access_configurations | json | X | √ |  | 
| topics | string_array | X | √ |  | 
| region | string | X | √ |  | 
| filter_criteria | json | X | √ |  | 
| function_response_types | string_array | X | √ |  | 
| self_managed_event_source | json | X | √ |  | 
| state | string | X | √ |  | 
| destination_config | json | X | √ |  | 
| tumbling_window_in_seconds | int | X | √ |  | 
| maximum_batching_window_in_seconds | int | X | √ |  | 
| maximum_record_age_in_seconds | int | X | √ |  | 
| maximum_retry_attempts | int | X | √ |  | 
| state_transition_reason | string | X | √ |  | 
| uuid | string | X | √ |  | 
| function_arn | string | X | √ |  | 
| self_managed_kafka_event_source_config | json | X | √ |  | 
| selefra_id | string | √ | √ | random id | 
| last_processing_result | string | X | √ |  | 


