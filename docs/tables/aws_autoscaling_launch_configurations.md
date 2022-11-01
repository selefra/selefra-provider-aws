# Table: aws_autoscaling_launch_configurations

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| iam_instance_profile | string | X | √ |  | 
| kernel_id | string | X | √ |  | 
| security_groups | string_array | X | √ |  | 
| spot_price | string | X | √ |  | 
| account_id | string | X | √ |  | 
| instance_type | string | X | √ |  | 
| launch_configuration_name | string | X | √ |  | 
| instance_monitoring | json | X | √ |  | 
| key_name | string | X | √ |  | 
| associate_public_ip_address | bool | X | √ |  | 
| block_device_mappings | json | X | √ |  | 
| ebs_optimized | bool | X | √ |  | 
| classic_link_vpc_security_groups | string_array | X | √ |  | 
| ramdisk_id | string | X | √ |  | 
| user_data | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| region | string | X | √ |  | 
| arn | string | √ | √ |  | 
| classic_link_vpc_id | string | X | √ |  | 
| placement_tenancy | string | X | √ |  | 
| created_time | timestamp | X | √ |  | 
| image_id | string | X | √ |  | 
| metadata_options | json | X | √ |  | 


