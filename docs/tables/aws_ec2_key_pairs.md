# Table: aws_ec2_key_pairs

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| create_time | timestamp | X | √ |  | 
| public_key | string | X | √ |  | 
| tags | json | X | √ |  | 
| key_pair_id | string | X | √ |  | 
| key_type | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| arn | string | √ | √ |  | 
| key_fingerprint | string | X | √ |  | 
| key_name | string | X | √ |  | 


