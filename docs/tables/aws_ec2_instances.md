# Table: aws_ec2_instances

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| ami_launch_index | int | X | √ |  | 
| block_device_mappings | json | X | √ |  | 
| licenses | json | X | √ |  | 
| private_dns_name_options | json | X | √ |  | 
| public_ip_address | string | X | √ |  | 
| arn | string | √ | √ |  | 
| architecture | string | X | √ |  | 
| capacity_reservation_id | string | X | √ |  | 
| instance_type | string | X | √ |  | 
| state | json | X | √ |  | 
| elastic_gpu_associations | json | X | √ |  | 
| monitoring | json | X | √ |  | 
| tags | json | X | √ |  | 
| hypervisor | string | X | √ |  | 
| instance_lifecycle | string | X | √ |  | 
| key_name | string | X | √ |  | 
| platform_details | string | X | √ |  | 
| root_device_type | string | X | √ |  | 
| instance_id | string | X | √ |  | 
| ipv6_address | string | X | √ |  | 
| kernel_id | string | X | √ |  | 
| platform | string | X | √ |  | 
| capacity_reservation_specification | json | X | √ |  | 
| image_id | string | X | √ |  | 
| private_dns_name | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| cpu_options | json | X | √ |  | 
| iam_instance_profile | json | X | √ |  | 
| launch_time | timestamp | X | √ |  | 
| security_groups | json | X | √ |  | 
| elastic_inference_accelerator_associations | json | X | √ |  | 
| enclave_options | json | X | √ |  | 
| ramdisk_id | string | X | √ |  | 
| state_reason | json | X | √ |  | 
| hibernation_options | json | X | √ |  | 
| private_ip_address | string | X | √ |  | 
| usage_operation_update_time | timestamp | X | √ |  | 
| vpc_id | string | X | √ |  | 
| root_device_name | string | X | √ |  | 
| client_token | string | X | √ |  | 
| outpost_arn | string | X | √ |  | 
| placement | json | X | √ |  | 
| product_codes | json | X | √ |  | 
| state_transition_reason | string | X | √ |  | 
| boot_mode | string | X | √ |  | 
| public_dns_name | string | X | √ |  | 
| spot_instance_request_id | string | X | √ |  | 
| maintenance_options | json | X | √ |  | 
| source_dest_check | bool | X | √ |  | 
| tpm_support | string | X | √ |  | 
| usage_operation | string | X | √ |  | 
| virtualization_type | string | X | √ |  | 
| account_id | string | X | √ |  | 
| state_transition_reason_time | timestamp | X | √ |  | 
| ebs_optimized | bool | X | √ |  | 
| ena_support | bool | X | √ |  | 
| metadata_options | json | X | √ |  | 
| region | string | X | √ |  | 
| network_interfaces | json | X | √ |  | 
| sriov_net_support | string | X | √ |  | 
| subnet_id | string | X | √ |  | 


