# Table: snowflake_account_parameter

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| default | string | X | √ | Default value of the parameter. | 
| description | string | X | √ | Description for the parameter. | 
| level | string | X | √ | Level of the parameter. Can be SYSTEM or ACCOUNT. | 
| type | string | X | √ | Data type of the parameter value. | 
| region | string | X | √ | The Snowflake region in which the account is located. | 
| account | string | X | √ | The Snowflake account ID. | 
| key | string | X | √ | Name of the account parameter. | 
| value | string | X | √ | Current value of the parameter. | 


