# Table: aws_guardduty_detector_members

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| administrator_id | string | X | √ |  | 
| detector_id | string | X | √ |  | 
| region | string | X | √ |  | 
| account_id | string | X | √ |  | 
| master_id | string | X | √ |  | 
| relationship_status | string | X | √ |  | 
| selefra_id | string | √ | √ | random id | 
| aws_guardduty_detectors_selefra_id | string | X | X | fk to aws_guardduty_detectors.selefra_id | 
| detector_arn | string | X | √ |  | 
| email | string | X | √ |  | 
| updated_at | string | X | √ |  | 
| invited_at | string | X | √ |  | 


