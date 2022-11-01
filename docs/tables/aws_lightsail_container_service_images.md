# Table: aws_lightsail_container_service_images

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| created_at | timestamp | X | √ |  | 
| digest | string | X | √ |  | 
| image | string | X | √ |  | 
| selefra_id | string | √ | √ | random id | 
| aws_lightsail_container_services_selefra_id | string | X | X | fk to aws_lightsail_container_services.selefra_id | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| container_service_arn | string | X | √ |  | 


