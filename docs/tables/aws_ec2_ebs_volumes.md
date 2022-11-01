# Table: aws_ec2_ebs_volumes

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| arn | string | √ | √ |  | 
| outpost_arn | string | X | √ |  | 
| snapshot_id | string | X | √ |  | 
| state | string | X | √ |  | 
| size | int | X | √ |  | 
| tags | json | X | √ |  | 
| throughput | int | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| attachments | json | X | √ |  | 
| create_time | timestamp | X | √ |  | 
| kms_key_id | string | X | √ |  | 
| volume_type | string | X | √ |  | 
| availability_zone | string | X | √ |  | 
| encrypted | bool | X | √ |  | 
| fast_restored | bool | X | √ |  | 
| iops | int | X | √ |  | 
| multi_attach_enabled | bool | X | √ |  | 
| volume_id | string | X | √ |  | 


