# Table: aws_iam_group_policies

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| policy_document | json | X | √ |  | 
| group_name | string | X | √ |  | 
| policy_name | string | X | √ |  | 
| aws_iam_groups_selefra_id | string | X | X | fk to aws_iam_groups.selefra_id | 
| group_arn | string | X | √ |  | 
| group_id | string | X | √ |  | 
| selefra_id | string | √ | √ | random id | 
| account_id | string | X | √ |  | 
| result_metadata | json | X | √ |  | 


