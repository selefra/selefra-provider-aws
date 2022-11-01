# Table: aws_emr_clusters

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| step_concurrency_level | int | X | √ |  | 
| auto_terminate | bool | X | √ |  | 
| custom_ami_id | string | X | √ |  | 
| outpost_arn | string | X | √ |  | 
| status | json | X | √ |  | 
| arn | string | √ | √ |  | 
| tags | json | X | √ |  | 
| normalized_instance_hours | int | X | √ |  | 
| requested_ami_version | string | X | √ |  | 
| security_configuration | string | X | √ |  | 
| service_role | string | X | √ |  | 
| termination_protected | bool | X | √ |  | 
| kerberos_attributes | json | X | √ |  | 
| log_encryption_kms_key_id | string | X | √ |  | 
| name | string | X | √ |  | 
| repo_upgrade_on_boot | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| configurations | json | X | √ |  | 
| instance_collection_type | string | X | √ |  | 
| release_label | string | X | √ |  | 
| scale_down_behavior | string | X | √ |  | 
| visible_to_all_users | bool | X | √ |  | 
| account_id | string | X | √ |  | 
| auto_scaling_role | string | X | √ |  | 
| id | string | X | √ |  | 
| log_uri | string | X | √ |  | 
| region | string | X | √ |  | 
| ebs_root_volume_size | int | X | √ |  | 
| master_public_dns_name | string | X | √ |  | 
| os_release_label | string | X | √ |  | 
| placement_groups | json | X | √ |  | 
| running_ami_version | string | X | √ |  | 
| applications | json | X | √ |  | 
| ec2_instance_attributes | json | X | √ |  | 


