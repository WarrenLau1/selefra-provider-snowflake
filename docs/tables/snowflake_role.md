# Table: snowflake_role

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| comment | string | X | √ | Comment for the role. | 
| granted_to_roles | int | X | √ | Number of roles that inherit the privileges of this role. | 
| is_default | string | X | √ | "Y" if is the default role of authenticated user, otherwise "F". | 
| is_inherited | string | X | √ | "Y" if current role is inherited by authenticated user, otherwise "F". | 
| owner | string | X | √ | Owner of the role. | 
| account | string | X | √ | The Snowflake account ID. | 
| name | string | X | √ | Name of the role. | 
| assigned_to_users | int | X | √ | Number of users the role is assigned. | 
| granted_roles | int | X | √ | Number of roles inherited by this role. | 
| is_current | string | X | √ | "Y" if is the current role of authenticated user, otherwise "F". | 
| region | string | X | √ | The Snowflake region in which the account is located. | 
| created_on | timestamp | X | √ | Date and time when the role was created. | 


