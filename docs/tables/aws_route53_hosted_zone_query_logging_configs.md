# Table: aws_route53_hosted_zone_query_logging_configs

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| arn | string | √ | √ |  | 
| hosted_zone_arn | string | X | √ |  | 
| cloud_watch_logs_log_group_arn | string | X | √ |  | 
| hosted_zone_id | string | X | √ |  | 
| id | string | X | √ |  | 
| aws_route53_hosted_zones_selefra_id | string | X | X | fk to aws_route53_hosted_zones.selefra_id | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| account_id | string | X | √ |  | 


