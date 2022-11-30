# Table: snowflake_view_grant

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| grantee_name | string | X | √ | Name of the object role has been granted. | 
| region | string | X | √ | The Snowflake region in which the account is located. | 
| privilege | string | X | √ | A defined level of access to an object. | 
| created_on | timestamp | X | √ | Date and time privilege was granted. | 
| grant_option | bool | X | √ | If set to TRUE, the recipient role can grant the privilege to other roles. | 
| granted_by | string | X | √ | Name of the object that granted access on the role. | 
| granted_on | string | X | √ | Date and time when the access was granted. | 
| granted_to | string | X | √ | Type of the object. | 
| account | string | X | √ | The Snowflake account ID. | 


