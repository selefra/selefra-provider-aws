# Table: aws_elasticache_reserved_cache_nodes

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| fixed_price | float | X | √ |  | 
| usage_price | float | X | √ |  | 
| duration | int | X | √ |  | 
| product_description | string | X | √ |  | 
| start_time | timestamp | X | √ |  | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| arn | string | √ | √ |  | 
| cache_node_count | int | X | √ |  | 
| offering_type | string | X | √ |  | 
| reserved_cache_node_id | string | X | √ |  | 
| reserved_cache_nodes_offering_id | string | X | √ |  | 
| cache_node_type | string | X | √ |  | 
| recurring_charges | json | X | √ |  | 
| state | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 


