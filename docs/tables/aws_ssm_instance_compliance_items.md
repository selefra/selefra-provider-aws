# Table: aws_ssm_instance_compliance_items

## Primary Keys 

```
id, instance_arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| details | json | X | √ |  | 
| resource_id | string | X | √ |  | 
| severity | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| status | string | X | √ |  | 
| title | string | X | √ |  | 
| aws_ssm_instances_selefra_id | string | X | X | fk to aws_ssm_instances.selefra_id | 
| compliance_type | string | X | √ |  | 
| resource_type | string | X | √ |  | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| id | string | X | √ |  | 
| instance_arn | string | X | √ |  | 
| execution_summary | json | X | √ |  | 


