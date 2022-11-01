# Table: aws_qldb_ledger_journal_s3_exports

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| ledger_arn | string | X | √ |  | 
| export_creation_time | timestamp | X | √ |  | 
| aws_qldb_ledgers_selefra_id | string | X | X | fk to aws_qldb_ledgers.selefra_id | 
| account_id | string | X | √ |  | 
| export_id | string | X | √ |  | 
| role_arn | string | X | √ |  | 
| s3_export_configuration | json | X | √ |  | 
| status | string | X | √ |  | 
| inclusive_start_time | timestamp | X | √ |  | 
| selefra_id | string | √ | √ | random id | 
| region | string | X | √ |  | 
| exclusive_end_time | timestamp | X | √ |  | 
| ledger_name | string | X | √ |  | 
| output_format | string | X | √ |  | 


