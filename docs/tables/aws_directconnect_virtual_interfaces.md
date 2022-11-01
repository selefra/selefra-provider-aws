# Table: aws_directconnect_virtual_interfaces

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| tags | json | X | √ |  | 
| mtu | int | X | √ |  | 
| virtual_interface_state | string | X | √ |  | 
| route_filter_prefixes | json | X | √ |  | 
| arn | string | √ | √ |  | 
| id | string | X | √ |  | 
| auth_key | string | X | √ |  | 
| connection_id | string | X | √ |  | 
| customer_address | string | X | √ |  | 
| jumbo_frame_capable | bool | X | √ |  | 
| owner_account | string | X | √ |  | 
| region | string | X | √ |  | 
| amazon_address | string | X | √ |  | 
| asn | int | X | √ |  | 
| virtual_interface_type | string | X | √ |  | 
| virtual_interface_name | string | X | √ |  | 
| vlan | int | X | √ |  | 
| account_id | string | X | √ |  | 
| address_family | string | X | √ |  | 
| aws_logical_device_id | string | X | √ |  | 
| location | string | X | √ |  | 
| amazon_side_asn | int | X | √ |  | 
| aws_device_v2 | string | X | √ |  | 
| bgp_peers | json | X | √ |  | 
| virtual_gateway_id | string | X | √ |  | 
| direct_connect_gateway_id | string | X | √ |  | 
| customer_router_config | string | X | √ |  | 
| site_link_enabled | bool | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 


