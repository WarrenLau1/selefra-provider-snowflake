# Table: snowflake_session_policy

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| created_on | timestamp | X | √ | Date and time of the creation of session policy. | 
| schema_name | string | X | √ | Name of the schema in database policy belongs. | 
| comment | string | X | √ | Comment for this policy. | 
| region | string | X | √ | The Snowflake region in which the account is located. | 
| name | string | X | √ | Identifier for the session policy. | 
| database_name | string | X | √ | Name of the database policy belongs. | 
| kind | string | X | √ | Type of the snowflake policy. | 
| owner | string | X | √ | Name of the role that owns the policy. | 
| session_idle_timeout_mins | int | X | √ | Time period in minutes of inactivity with either the web interface or a programmatic client. | 
| session_ui_idle_timeout_mins | int | X | √ | Time period in minutes of inactivity with the web interface. | 
| account | string | X | √ | The Snowflake account ID. | 


