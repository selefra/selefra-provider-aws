# Table: aws_ec2_instance_statuses

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| instance_status | json | X | √ |  | 
| outpost_arn | string | X | √ |  | 
| system_status | json | X | √ |  | 
| account_id | string | X | √ |  | 
| instance_id | string | X | √ |  | 
| instance_state | json | X | √ |  | 
| events | json | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| region | string | X | √ |  | 
| arn | string | √ | √ |  | 
| availability_zone | string | X | √ |  | 


