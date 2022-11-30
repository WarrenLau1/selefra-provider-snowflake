# Table: snowflake_login_history

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| account | string | X | √ | The Snowflake account ID. | 
| reported_client_version | string | X | √ | Reported version of the client software. This information is not authenticated. | 
| second_authentication_factor | string | X | √ | The second factor, if using multi factor authentication, or NULL otherwise. | 
| is_success | string | X | √ | Whether the user's request was successful or not. | 
| error_code | int | X | √ | Error code, if the request was not successful. | 
| event_id | string | X | √ | Internal/system-generated identifier for the login attempt. | 
| client_ip | string | X | √ | IP address where the request originated from. | 
| first_authentication_factor | string | X | √ | Method used to authenticate the user (the first factor, if using multi factor authentication). | 
| region | string | X | √ | The Snowflake region in which the account is located. | 
| user_name | string | X | √ | User associated with this event. | 
| event_type | string | X | √ | Event type, such as LOGIN for authentication events. | 
| reported_client_type | string | X | √ | Reported type of the client software, such as JDBC_DRIVER, ODBC_DRIVER, etc. This information is not authenticated. | 
| error_message | string | X | √ | Error message returned to the user, if the request was not successful. | 
| related_event_id | int | X | √ | Reserved for future use. | 
| event_timestamp | timestamp | X | √ | Time (in the UTC time zone) of the event occurrence. | 


