# Table: aws_guardduty_detectors

## Primary Keys 

```
account_id, region, id
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| service_role | string | X | √ |  | 
| created_at | string | X | √ |  | 
| data_sources | json | X | √ |  | 
| finding_publishing_frequency | string | X | √ |  | 
| tags | json | X | √ |  | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| id | string | X | √ |  | 
| updated_at | string | X | √ |  | 
| result_metadata | json | X | √ |  | 
| arn | string | X | √ |  | 
| status | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 


