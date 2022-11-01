# Table: aws_wafregional_web_acls

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| name | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| region | string | X | √ |  | 
| resources_for_web_acl | string_array | X | √ |  | 
| default_action | json | X | √ |  | 
| rules | json | X | √ |  | 
| web_acl_id | string | X | √ |  | 
| metric_name | string | X | √ |  | 
| account_id | string | X | √ |  | 
| arn | string | √ | √ |  | 
| tags | json | X | √ | `Web ACL tags.` | 


