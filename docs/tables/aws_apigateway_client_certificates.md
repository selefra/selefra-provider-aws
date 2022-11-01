# Table: aws_apigateway_client_certificates

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| arn | string | √ | √ |  | 
| created_date | timestamp | X | √ |  | 
| pem_encoded_certificate | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| expiration_date | timestamp | X | √ |  | 
| tags | json | X | √ |  | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| client_certificate_id | string | X | √ |  | 
| description | string | X | √ |  | 


