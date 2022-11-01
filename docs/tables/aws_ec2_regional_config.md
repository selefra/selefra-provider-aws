# Table: aws_ec2_regional_config

## Primary Keys 

```
account_id, region
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| ebs_encryption_enabled_by_default | bool | X | √ |  | 
| ebs_default_kms_key_id | string | X | √ |  | 


