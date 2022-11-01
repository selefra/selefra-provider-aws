# Table: aws_elasticbeanstalk_applications

## Primary Keys 

```
arn, date_created
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| date_updated | timestamp | X | √ |  | 
| description | string | X | √ |  | 
| versions | string_array | X | √ |  | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| date_created | timestamp | X | √ |  | 
| configuration_templates | string_array | X | √ |  | 
| arn | string | X | √ |  | 
| application_name | string | X | √ |  | 
| resource_lifecycle_config | json | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 


