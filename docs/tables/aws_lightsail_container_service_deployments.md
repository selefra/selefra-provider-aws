# Table: aws_lightsail_container_service_deployments

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| aws_lightsail_container_services_selefra_id | string | X | X | fk to aws_lightsail_container_services.selefra_id | 
| region | string | X | √ |  | 
| container_service_arn | string | X | √ |  | 
| containers | json | X | √ |  | 
| state | string | X | √ |  | 
| selefra_id | string | √ | √ | random id | 
| account_id | string | X | √ |  | 
| created_at | timestamp | X | √ |  | 
| public_endpoint | json | X | √ |  | 
| version | int | X | √ |  | 


