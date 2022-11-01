# Table: aws_ecr_repositories

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| repository_uri | string | X | √ |  | 
| account_id | string | X | √ |  | 
| arn | string | √ | √ |  | 
| tags | json | X | √ |  | 
| image_tag_mutability | string | X | √ |  | 
| registry_id | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| region | string | X | √ |  | 
| created_at | timestamp | X | √ |  | 
| encryption_configuration | json | X | √ |  | 
| image_scanning_configuration | json | X | √ |  | 
| repository_name | string | X | √ |  | 


