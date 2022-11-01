# Table: aws_route53_hosted_zone_resource_record_sets

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| account_id | string | X | √ |  | 
| resource_records | json | X | √ |  | 
| traffic_policy_instance_id | string | X | √ |  | 
| selefra_id | string | √ | √ | random id | 
| cidr_routing_config | json | X | √ |  | 
| failover | string | X | √ |  | 
| multi_value_answer | bool | X | √ |  | 
| set_identifier | string | X | √ |  | 
| aws_route53_hosted_zones_selefra_id | string | X | X | fk to aws_route53_hosted_zones.selefra_id | 
| hosted_zone_arn | string | X | √ |  | 
| name | string | X | √ |  | 
| geo_location | json | X | √ |  | 
| health_check_id | string | X | √ |  | 
| region | string | X | √ |  | 
| type | string | X | √ |  | 
| alias_target | json | X | √ |  | 
| ttl | int | X | √ |  | 
| weight | int | X | √ |  | 


