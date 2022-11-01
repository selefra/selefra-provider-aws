# Table: aws_ssm_documents

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| review_status | string | X | √ |  | 
| version_name | string | X | √ |  | 
| document_type | string | X | √ |  | 
| hash_type | string | X | √ |  | 
| pending_review_version | string | X | √ |  | 
| arn | string | √ | √ |  | 
| created_date | timestamp | X | √ |  | 
| document_format | string | X | √ |  | 
| default_version | string | X | √ |  | 
| description | string | X | √ |  | 
| display_name | string | X | √ |  | 
| schema_version | string | X | √ |  | 
| region | string | X | √ |  | 
| tags | json | X | √ |  | 
| approved_version | string | X | √ |  | 
| author | string | X | √ |  | 
| hash | string | X | √ |  | 
| sha1 | string | X | √ |  | 
| owner | string | X | √ |  | 
| status | string | X | √ |  | 
| account_id | string | X | √ |  | 
| category | string_array | X | √ |  | 
| document_version | string | X | √ |  | 
| permissions | json | X | √ |  | 
| latest_version | string | X | √ |  | 
| requires | json | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| parameters | json | X | √ |  | 
| review_information | json | X | √ |  | 
| target_type | string | X | √ |  | 
| platform_types | string_array | X | √ |  | 
| status_information | string | X | √ |  | 
| attachments_information | json | X | √ |  | 
| category_enum | string_array | X | √ |  | 
| name | string | X | √ |  | 


