# Table: aws_ec2_vpc_endpoint_service_configurations

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| arn | string | √ | √ |  | 
| service_id | string | X | √ |  | 
| service_name | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| acceptance_required | bool | X | √ |  | 
| gateway_load_balancer_arns | string_array | X | √ |  | 
| manages_vpc_endpoints | bool | X | √ |  | 
| network_load_balancer_arns | string_array | X | √ |  | 
| private_dns_name | string | X | √ |  | 
| tags | json | X | √ |  | 
| account_id | string | X | √ |  | 
| base_endpoint_dns_names | string_array | X | √ |  | 
| private_dns_name_configuration | json | X | √ |  | 
| service_type | json | X | √ |  | 
| region | string | X | √ |  | 
| availability_zones | string_array | X | √ |  | 
| payer_responsibility | string | X | √ |  | 
| service_state | string | X | √ |  | 
| supported_ip_address_types | string_array | X | √ |  | 


