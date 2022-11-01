# Table: aws_route53_traffic_policy_versions

## Primary Keys 

```
traffic_policy_arn, id, version
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| type | string | X | √ |  | 
| comment | string | X | √ |  | 
| aws_route53_traffic_policies_selefra_id | string | X | X | fk to aws_route53_traffic_policies.selefra_id | 
| account_id | string | X | √ |  | 
| traffic_policy_arn | string | X | √ |  | 
| id | string | X | √ |  | 
| version | int | X | √ |  | 
| document | json | X | √ |  | 
| name | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 


