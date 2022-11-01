# Table: aws_ec2_byoip_cidrs

## Primary Keys 

```
account_id, region, cidr
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| state | string | X | √ |  | 
| status_message | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| cidr | string | X | √ |  | 
| description | string | X | √ |  | 


