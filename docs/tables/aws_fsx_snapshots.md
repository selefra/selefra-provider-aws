# Table: aws_fsx_snapshots

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| creation_time | timestamp | X | √ |  | 
| lifecycle | string | X | √ |  | 
| name | string | X | √ |  | 
| volume_id | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| lifecycle_transition_reason | json | X | √ |  | 
| snapshot_id | string | X | √ |  | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| arn | string | √ | √ |  | 
| tags | json | X | √ |  | 
| administrative_actions | json | X | √ |  | 


