# Table: snowflake_role_grant

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| role | string | X | √ | Name of the role on that access has been granted. | 
| created_on | timestamp | X | √ | Date and time when the role was granted to the user/role. | 
| granted_to | string | X | √ | Type of the object. Valid values USER and ROLE. | 
| grantee_name | string | X | √ | Name of the object role has been granted. | 
| granted_by | string | X | √ | Name of the object that granted access on the role. | 
| region | string | X | √ | The Snowflake region in which the account is located. | 
| account | string | X | √ | The Snowflake account ID. | 


