# Table: aws_config_conformance_pack_rule_compliances

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| config_rule_name | string | X | √ |  | 
| controls | string_array | X | √ |  | 
| evaluation_result_identifier | json | X | √ |  | 
| result_recorded_time | timestamp | X | √ |  | 
| annotation | string | X | √ |  | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| compliance_type | string | X | √ |  | 
| selefra_id | string | √ | √ | random id | 
| conformance_pack_arn | string | X | √ |  | 
| config_rule_invoked_time | timestamp | X | √ |  | 
| aws_config_conformance_packs_selefra_id | string | X | X | fk to aws_config_conformance_packs.selefra_id | 


