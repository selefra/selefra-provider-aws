# Table: aws_shield_attacks

## Primary Keys 

```
id
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| attack_counters | json | X | √ |  | 
| attack_properties | json | X | √ |  | 
| mitigations | json | X | √ |  | 
| start_time | timestamp | X | √ |  | 
| account_id | string | X | √ |  | 
| id | string | √ | √ | `The unique identifier (ID) of the attack` | 
| end_time | timestamp | X | √ |  | 
| resource_arn | string | X | √ |  | 
| sub_resources | json | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 


