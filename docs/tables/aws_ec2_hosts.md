# Table: aws_ec2_hosts

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| arn | string | √ | √ |  | 
| availability_zone_id | string | X | √ |  | 
| release_time | timestamp | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| outpost_arn | string | X | √ |  | 
| tags | json | X | √ |  | 
| allocation_time | timestamp | X | √ |  | 
| available_capacity | json | X | √ |  | 
| client_token | string | X | √ |  | 
| instances | json | X | √ |  | 
| auto_placement | string | X | √ |  | 
| host_id | string | X | √ |  | 
| member_of_service_linked_resource_group | bool | X | √ |  | 
| host_properties | json | X | √ |  | 
| host_recovery | string | X | √ |  | 
| host_reservation_id | string | X | √ |  | 
| owner_id | string | X | √ |  | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| allows_multiple_instance_types | string | X | √ |  | 
| availability_zone | string | X | √ |  | 
| state | string | X | √ |  | 


