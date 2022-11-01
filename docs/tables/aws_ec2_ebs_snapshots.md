# Table: aws_ec2_ebs_snapshots

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| region | string | X | √ |  | 
| outpost_arn | string | X | √ |  | 
| owner_id | string | X | √ |  | 
| kms_key_id | string | X | √ |  | 
| progress | string | X | √ |  | 
| snapshot_id | string | X | √ |  | 
| start_time | timestamp | X | √ |  | 
| state_message | string | X | √ |  | 
| storage_tier | string | X | √ |  | 
| volume_size | int | X | √ |  | 
| arn | string | √ | √ |  | 
| attribute | json | X | √ |  | 
| restore_expiry_time | timestamp | X | √ |  | 
| encrypted | bool | X | √ |  | 
| owner_alias | string | X | √ |  | 
| state | string | X | √ |  | 
| tags | json | X | √ |  | 
| volume_id | string | X | √ |  | 
| account_id | string | X | √ |  | 
| data_encryption_key_id | string | X | √ |  | 
| description | string | X | √ |  | 


