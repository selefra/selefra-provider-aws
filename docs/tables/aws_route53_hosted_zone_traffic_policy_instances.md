# Table: aws_route53_hosted_zone_traffic_policy_instances

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| aws_route53_hosted_zones_selefra_id | string | X | X | fk to aws_route53_hosted_zones.selefra_id | 
| hosted_zone_id | string | X | √ |  | 
| id | string | X | √ |  | 
| name | string | X | √ |  | 
| traffic_policy_type | string | X | √ |  | 
| traffic_policy_id | string | X | √ |  | 
| traffic_policy_version | int | X | √ |  | 
| account_id | string | X | √ |  | 
| arn | string | X | √ | `Amazon Resource Name (ARN) of the route53 hosted zone traffic policy instance.` | 
| message | string | X | √ |  | 
| ttl | int | X | √ |  | 
| hosted_zone_arn | string | X | √ |  | 
| state | string | X | √ |  | 
| selefra_id | string | √ | √ | random id | 


