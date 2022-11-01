# Table: aws_ec2_instance_types

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| hibernation_supported | bool | X | √ |  | 
| hypervisor | string | X | √ |  | 
| placement_group_info | json | X | √ |  | 
| processor_info | json | X | √ |  | 
| region | string | X | √ |  | 
| auto_recovery_supported | bool | X | √ |  | 
| inference_accelerator_info | json | X | √ |  | 
| memory_info | json | X | √ |  | 
| network_info | json | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| account_id | string | X | √ |  | 
| bare_metal | bool | X | √ |  | 
| current_generation | bool | X | √ |  | 
| fpga_info | json | X | √ |  | 
| supported_boot_modes | string_array | X | √ |  | 
| supported_root_device_types | string_array | X | √ |  | 
| supported_usage_classes | string_array | X | √ |  | 
| burstable_performance_supported | bool | X | √ |  | 
| gpu_info | json | X | √ |  | 
| instance_type | string | X | √ |  | 
| v_cpu_info | json | X | √ |  | 
| free_tier_eligible | bool | X | √ |  | 
| instance_storage_info | json | X | √ |  | 
| supported_virtualization_types | string_array | X | √ |  | 
| instance_storage_supported | bool | X | √ |  | 
| arn | string | √ | √ |  | 
| dedicated_hosts_supported | bool | X | √ |  | 
| ebs_info | json | X | √ |  | 


