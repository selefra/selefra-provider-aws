# Table: aws_eks_clusters

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| client_request_token | string | X | √ |  | 
| kubernetes_network_config | json | X | √ |  | 
| version | string | X | √ |  | 
| region | string | X | √ |  | 
| certificate_authority | json | X | √ |  | 
| encryption_config | json | X | √ |  | 
| outpost_config | json | X | √ |  | 
| status | string | X | √ |  | 
| arn | string | √ | √ |  | 
| endpoint | string | X | √ |  | 
| health | json | X | √ |  | 
| id | string | X | √ |  | 
| logging | json | X | √ |  | 
| platform_version | string | X | √ |  | 
| resources_vpc_config | json | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| account_id | string | X | √ |  | 
| connector_config | json | X | √ |  | 
| created_at | timestamp | X | √ |  | 
| identity | json | X | √ |  | 
| name | string | X | √ |  | 
| role_arn | string | X | √ |  | 
| tags | json | X | √ |  | 


