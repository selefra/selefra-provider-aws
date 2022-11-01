# Table: aws_lightsail_load_balancers

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| name | string | X | √ |  | 
| resource_type | string | X | √ |  | 
| support_code | string | X | √ |  | 
| tags | json | X | √ |  | 
| configuration_options | json | X | √ |  | 
| created_at | timestamp | X | √ |  | 
| https_redirection_enabled | bool | X | √ |  | 
| instance_health_summary | json | X | √ |  | 
| public_ports | int_array | X | √ |  | 
| health_check_path | string | X | √ |  | 
| ip_address_type | string | X | √ |  | 
| state | string | X | √ |  | 
| tls_certificate_summaries | json | X | √ |  | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| arn | string | √ | √ |  | 
| dns_name | string | X | √ |  | 
| instance_port | int | X | √ |  | 
| location | json | X | √ |  | 
| protocol | string | X | √ |  | 
| tls_policy_name | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 


