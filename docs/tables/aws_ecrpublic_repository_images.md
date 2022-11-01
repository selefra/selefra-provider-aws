# Table: aws_ecrpublic_repository_images

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| image_digest | string | X | √ |  | 
| image_pushed_at | timestamp | X | √ |  | 
| aws_ecrpublic_repositories_selefra_id | string | X | X | fk to aws_ecrpublic_repositories.selefra_id | 
| artifact_media_type | string | X | √ |  | 
| region | string | X | √ |  | 
| arn | string | √ | √ |  | 
| image_manifest_media_type | string | X | √ |  | 
| image_size_in_bytes | int | X | √ |  | 
| image_tags | string_array | X | √ |  | 
| registry_id | string | X | √ |  | 
| repository_name | string | X | √ |  | 
| account_id | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 


