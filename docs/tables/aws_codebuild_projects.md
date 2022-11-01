# Table: aws_codebuild_projects

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| logs_config | json | X | √ |  | 
| resource_access_role | string | X | √ |  | 
| service_role | string | X | √ |  | 
| vpc_config | json | X | √ |  | 
| arn | string | √ | √ |  | 
| artifacts | json | X | √ |  | 
| concurrent_build_limit | int | X | √ |  | 
| build_batch_config | json | X | √ |  | 
| tags | json | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| created | timestamp | X | √ |  | 
| description | string | X | √ |  | 
| last_modified | timestamp | X | √ |  | 
| secondary_source_versions | json | X | √ |  | 
| account_id | string | X | √ |  | 
| environment | json | X | √ |  | 
| file_system_locations | json | X | √ |  | 
| project_visibility | string | X | √ |  | 
| source_version | string | X | √ |  | 
| secondary_artifacts | json | X | √ |  | 
| region | string | X | √ |  | 
| cache | json | X | √ |  | 
| encryption_key | string | X | √ |  | 
| badge | json | X | √ |  | 
| source | json | X | √ |  | 
| webhook | json | X | √ |  | 
| secondary_sources | json | X | √ |  | 
| timeout_in_minutes | int | X | √ |  | 
| name | string | X | √ |  | 
| public_project_alias | string | X | √ |  | 
| queued_timeout_in_minutes | int | X | √ |  | 


