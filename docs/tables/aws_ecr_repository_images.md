# Table: aws_ecr_repository_images

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| image_digest | string | X | √ |  | 
| image_tags | string_array | X | √ |  | 
| last_recorded_pull_time | timestamp | X | √ |  | 
| aws_ecr_repositories_selefra_id | string | X | X | fk to aws_ecr_repositories.selefra_id | 
| image_pushed_at | timestamp | X | √ |  | 
| image_size_in_bytes | int | X | √ |  | 
| repository_name | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| image_scan_findings_summary | json | X | √ |  | 
| registry_id | string | X | √ |  | 
| arn | string | √ | √ |  | 
| artifact_media_type | string | X | √ |  | 
| image_manifest_media_type | string | X | √ |  | 
| image_scan_status | json | X | √ |  | 


