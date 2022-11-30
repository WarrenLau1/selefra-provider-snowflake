# Table: snowflake_session

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| authentication_method | string | X | √ | The authentication method used to access Snowflake. | 
| client_application_id | string | X | √ | The identifier for the Snowflake-provided client application used to create the remote session to Snowflake (e.g. JDBC 3.8.7) | 
| client_application_version | string | X | √ | The version number (e.g. 3.8.7) of the Snowflake-provided client application used to create the remote session to Snowflake. | 
| region | string | X | √ | The Snowflake region in which the account is located. | 
| account | string | X | √ | The Snowflake account ID. | 
| client_version | string | X | √ | The version number (e.g. 47154) of the third-party client application that uses a Snowflake-provided client to create a remote session to Snowflake, if available. | 
| login_event_id | string | X | √ | The unique identifier for the login event. | 
| session_id | string | X | √ | The unique identifier for the current session. | 
| user_name | string | X | √ | The user name of the user. | 
| created_on | timestamp | X | √ | Date and time when the session was created. | 
| client_build_id | string | X | √ | The build number (e.g. 41897) of the third-party client application used to create a remote session to Snowflake, if available. For example, a third-party Java application that uses the JDBC driver to connect to Snowflake. | 
| client_environment | json | X | √ | The environment variables (e.g. operating system, OCSP mode) of the client used to create a remote session to Snowflake. | 


