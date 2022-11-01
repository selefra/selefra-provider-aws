# Table: aws_elasticsearch_domains

## Primary Keys 

```
account_id, region, id
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| domain_name | string | X | √ |  | 
| encryption_at_rest_options | json | X | √ |  | 
| snapshot_options | json | X | √ |  | 
| upgrade_processing | bool | X | √ |  | 
| ebs_options | json | X | √ |  | 
| auto_tune_options | json | X | √ |  | 
| cognito_options | json | X | √ |  | 
| endpoints | json | X | √ |  | 
| region | string | X | √ |  | 
| created | bool | X | √ |  | 
| node_to_node_encryption_options | json | X | √ |  | 
| vpc_options | json | X | √ |  | 
| id | string | X | √ |  | 
| elasticsearch_cluster_config | json | X | √ |  | 
| access_policies | string | X | √ |  | 
| change_progress_details | json | X | √ |  | 
| tags | json | X | √ |  | 
| advanced_security_options | json | X | √ |  | 
| domain_endpoint_options | json | X | √ |  | 
| log_publishing_options | json | X | √ |  | 
| processing | bool | X | √ |  | 
| service_software_options | json | X | √ |  | 
| advanced_options | json | X | √ |  | 
| arn | string | X | √ |  | 
| deleted | bool | X | √ |  | 
| elasticsearch_version | string | X | √ |  | 
| endpoint | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| account_id | string | X | √ |  | 


