# Table: aws_ses_templates

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| region | string | X | √ |  | 
| arn | string | √ | √ |  | 
| template_name | string | X | √ |  | 
| text | string | X | √ |  | 
| account_id | string | X | √ |  | 
| html | string | X | √ |  | 
| subject | string | X | √ |  | 
| created_timestamp | timestamp | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 


