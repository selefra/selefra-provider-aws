# Table: aws_elasticache_reserved_cache_nodes_offerings

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| cache_node_type | string | X | √ |  | 
| recurring_charges | json | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| arn | string | √ | √ |  | 
| duration | int | X | √ |  | 
| fixed_price | float | X | √ |  | 
| offering_type | string | X | √ |  | 
| product_description | string | X | √ |  | 
| reserved_cache_nodes_offering_id | string | X | √ |  | 
| usage_price | float | X | √ |  | 


