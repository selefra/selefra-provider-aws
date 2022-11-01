# Table: aws_qldb_ledger_journal_kinesis_streams

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| selefra_id | string | √ | √ | random id | 
| kinesis_configuration | json | X | √ |  | 
| role_arn | string | X | √ |  | 
| stream_name | string | X | √ |  | 
| error_cause | string | X | √ |  | 
| ledger_name | string | X | √ |  | 
| exclusive_end_time | timestamp | X | √ |  | 
| inclusive_start_time | timestamp | X | √ |  | 
| ledger_arn | string | X | √ |  | 
| aws_qldb_ledgers_selefra_id | string | X | X | fk to aws_qldb_ledgers.selefra_id | 
| arn | string | X | √ |  | 
| creation_time | timestamp | X | √ |  | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| status | string | X | √ |  | 
| stream_id | string | X | √ |  | 


