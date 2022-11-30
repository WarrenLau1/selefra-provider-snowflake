# Table: snowflake_user_grant

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| account | string | X | √ | The Snowflake account ID. | 
| role | string | X | √ | Name of the role that has been granted to user. | 
| created_on | timestamp | X | √ | Date and time when the role was granted to the user/role. | 
| granted_to | string | X | √ | Type of the object. Only USER for this table. | 
| granted_by | string | X | √ | Name of the object that granted access on the user. | 
| region | string | X | √ | The Snowflake region in which the account is located. | 


