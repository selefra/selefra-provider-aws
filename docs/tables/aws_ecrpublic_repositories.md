# Table: aws_ecrpublic_repositories

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| arn | string | √ | √ |  | 
| tags | json | X | √ |  | 
| repository_uri | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| created_at | timestamp | X | √ |  | 
| registry_id | string | X | √ |  | 
| repository_name | string | X | √ |  | 


