# Table: snowflake_account_grant

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| granted_on | string | X | √ | Date and time when the access was granted. | 
| account | string | X | √ | The Snowflake account ID. | 
| privilege | string | X | √ | A defined level of access to an object. | 
| granted_by | string | X | √ | Name of the object that granted access on the role. | 
| grant_option | bool | X | √ | If set to TRUE, the recipient role can grant the privilege to other roles. | 
| granted_to | string | X | √ | Type of the object. | 
| grantee_name | string | X | √ | Name of the object role has been granted. | 
| region | string | X | √ | The Snowflake region in which the account is located. | 
| name | string | X | √ | An entity to which access can be granted. Unless allowed by a grant, access will be denied. | 
| created_on | timestamp | X | √ | Date and time privilege was granted. | 


