# Table: snowflake_database_grant

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| database | string | X | √ | Name of the database. | 
| grantee_name | string | X | √ | Name of the object role has been granted. | 
| account | string | X | √ | The Snowflake account ID. | 
| privilege | string | X | √ | A defined level of access to an database. | 
| created_on | timestamp | X | √ | Date and time when the access was granted. | 
| grant_option | bool | X | √ | If set to TRUE, the recipient role can grant the privilege to other roles. | 
| granted_by | string | X | √ | Identifier for the object that granted the privilege. | 
| granted_on | string | X | √ | Type of the object. | 
| granted_to | string | X | √ | Type of the object role has been granted. | 
| region | string | X | √ | The Snowflake region in which the account is located. | 


