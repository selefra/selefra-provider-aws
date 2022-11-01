# Table: aws_ec2_images

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| product_codes | json | X | √ |  | 
| ramdisk_id | string | X | √ |  | 
| sriov_net_support | string | X | √ |  | 
| state | string | X | √ |  | 
| image_location | string | X | √ |  | 
| image_type | string | X | √ |  | 
| kernel_id | string | X | √ |  | 
| state_reason | json | X | √ |  | 
| platform | string | X | √ |  | 
| block_device_mappings | json | X | √ |  | 
| ena_support | bool | X | √ |  | 
| name | string | X | √ |  | 
| platform_details | string | X | √ |  | 
| region | string | X | √ |  | 
| arn | string | √ | √ |  | 
| image_owner_alias | string | X | √ |  | 
| tpm_support | string | X | √ |  | 
| virtualization_type | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| architecture | string | X | √ |  | 
| boot_mode | string | X | √ |  | 
| image_id | string | X | √ |  | 
| root_device_type | string | X | √ |  | 
| creation_date | string | X | √ |  | 
| imds_support | string | X | √ |  | 
| root_device_name | string | X | √ |  | 
| description | string | X | √ |  | 
| hypervisor | string | X | √ |  | 
| public | bool | X | √ |  | 
| tags | json | X | √ |  | 
| usage_operation | string | X | √ |  | 
| account_id | string | X | √ |  | 
| deprecation_time | string | X | √ |  | 
| owner_id | string | X | √ |  | 


