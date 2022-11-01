# Table: aws_elasticbeanstalk_environments

## Primary Keys 

```
account_id, id
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| cname | string | X | √ |  | 
| date_updated | timestamp | X | √ |  | 
| endpoint_url | string | X | √ |  | 
| version_label | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| application_name | string | X | √ |  | 
| health | string | X | √ |  | 
| status | string | X | √ |  | 
| template_name | string | X | √ |  | 
| account_id | string | X | √ |  | 
| arn | string | X | √ |  | 
| abortable_operation_in_progress | bool | X | √ |  | 
| date_created | timestamp | X | √ |  | 
| description | string | X | √ |  | 
| operations_role | string | X | √ |  | 
| platform_arn | string | X | √ |  | 
| tier | json | X | √ |  | 
| resources | json | X | √ |  | 
| region | string | X | √ |  | 
| tags | json | X | √ |  | 
| id | string | X | √ |  | 
| listeners | json | X | √ |  | 
| environment_links | json | X | √ |  | 
| environment_name | string | X | √ |  | 
| health_status | string | X | √ |  | 
| solution_stack_name | string | X | √ |  | 


