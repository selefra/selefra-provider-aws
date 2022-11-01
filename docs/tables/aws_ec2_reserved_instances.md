# Table: aws_ec2_reserved_instances

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| arn | string | √ | √ |  | 
| currency_code | string | X | √ |  | 
| duration | int | X | √ |  | 
| account_id | string | X | √ |  | 
| instance_count | int | X | √ |  | 
| instance_tenancy | string | X | √ |  | 
| offering_class | string | X | √ |  | 
| offering_type | string | X | √ |  | 
| region | string | X | √ |  | 
| end | timestamp | X | √ |  | 
| fixed_price | float | X | √ |  | 
| product_description | string | X | √ |  | 
| scope | string | X | √ |  | 
| start | timestamp | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| availability_zone | string | X | √ |  | 
| instance_type | string | X | √ |  | 
| recurring_charges | json | X | √ |  | 
| reserved_instances_id | string | X | √ |  | 
| state | string | X | √ |  | 
| usage_price | float | X | √ |  | 
| tags | json | X | √ |  | 


