# Table: snowflake_network_policy

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| name | string | X | √ | Identifier for the network policy. | 
| comment | string | X | √ | Specifies a comment for the network policy. | 
| entries_in_blocked_ip_list | int | X | √ | No of entries in the blocked IP list. | 
| allowed_ip_list | string | X | √ | Comma-separated list of one or more IPv4 addresses that are allowed access to your Snowflake account. | 
| blocked_ip_list | string | X | √ | Comma-separated list of one or more IPv4 addresses that are denied access to your Snowflake account. | 
| region | string | X | √ | The Snowflake region in which the account is located. | 
| account | string | X | √ | The Snowflake account ID. | 
| created_on | timestamp | X | √ | Date and time when the policy was created. | 
| entries_in_allowed_ip_list | int | X | √ | No of entries in the allowed IP list. | 


