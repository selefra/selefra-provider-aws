# Table: aws_kms_keys

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| arn | string | √ | √ |  | 
| replica_keys | json | X | √ |  | 
| custom_key_store_id | string | X | √ |  | 
| key_spec | string | X | √ |  | 
| multi_region | bool | X | √ |  | 
| pending_deletion_window_in_days | int | X | √ |  | 
| expiration_model | string | X | √ |  | 
| key_manager | string | X | √ |  | 
| key_id | string | X | √ |  | 
| deletion_date | timestamp | X | √ |  | 
| signing_algorithms | string_array | X | √ |  | 
| valid_to | timestamp | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| cloud_hsm_cluster_id | string | X | √ |  | 
| creation_date | timestamp | X | √ |  | 
| description | string | X | √ |  | 
| tags | json | X | √ |  | 
| key_state | string | X | √ |  | 
| key_usage | string | X | √ |  | 
| mac_algorithms | string_array | X | √ |  | 
| rotation_enabled | bool | X | √ |  | 
| enabled | bool | X | √ |  | 
| encryption_algorithms | string_array | X | √ |  | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| multi_region_configuration | json | X | √ |  | 
| aws_account_id | string | X | √ |  | 
| customer_master_key_spec | string | X | √ |  | 
| origin | string | X | √ |  | 


