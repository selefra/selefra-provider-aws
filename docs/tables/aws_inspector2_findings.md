# Table: aws_inspector2_findings

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| region | string | X | √ |  | 
| last_observed_at | timestamp | X | √ |  | 
| type | string | X | √ |  | 
| first_observed_at | timestamp | X | √ |  | 
| remediation | json | X | √ |  | 
| status | string | X | √ |  | 
| inspector_score_details | json | X | √ |  | 
| updated_at | timestamp | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| account_id | string | X | √ |  | 
| arn | string | √ | √ |  | 
| description | string | X | √ |  | 
| inspector_score | float | X | √ |  | 
| network_reachability_details | json | X | √ |  | 
| title | string | X | √ |  | 
| aws_account_id | string | X | √ |  | 
| resources | json | X | √ |  | 
| severity | string | X | √ |  | 
| fix_available | string | X | √ |  | 
| package_vulnerability_details | json | X | √ |  | 


