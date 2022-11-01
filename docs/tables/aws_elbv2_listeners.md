# Table: aws_elbv2_listeners

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| account_id | string | X | √ |  | 
| tags | json | X | √ |  | 
| default_actions | json | X | √ |  | 
| port | int | X | √ |  | 
| ssl_policy | string | X | √ |  | 
| protocol | string | X | √ |  | 
| aws_elbv2_load_balancers_selefra_id | string | X | X | fk to aws_elbv2_load_balancers.selefra_id | 
| region | string | X | √ |  | 
| arn | string | √ | √ |  | 
| alpn_policy | string_array | X | √ |  | 
| certificates | json | X | √ |  | 
| load_balancer_arn | string | X | √ |  | 


