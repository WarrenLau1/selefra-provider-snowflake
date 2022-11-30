# Table: snowflake_user

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| name | string | X | √ | Name of the snowflake user. | 
| comment | string | X | √ | Comment associated to user in the dictionary. | 
| region | string | X | √ | The Snowflake region in which the account is located. | 
| custom_landing_page_url_flush_next_ui_load | bool | X | √ | The timestamp on which the last non-null password was set for the user. Default to null if no password has been set yet. | 
| default_role | string | X | √ | Primary principal of user session will be set to this role. | 
| display_name | string | X | √ | Display name of the user. | 
| first_name | string | X | √ | First name of the user. | 
| last_name | string | X | √ | Last name of the user. | 
| default_namespace | string | X | √ | Default database namespace prefix for this user. | 
| default_secondary_roles | string | X | √ | The secondary roles will be set to all roles provided here. | 
| mins_to_unlock | string | X | √ | Temporary lock on the user will be removed after specified number of minutes. | 
| snowflake_support | string | X | √ | Snowflake Support is allowed to use the user or account. | 
| has_password | bool | X | √ | Whether the user has password. | 
| disabled | string | X | √ | Whether the user is disabled. | 
| rsa_public_key_2_fp | string | X | √ | Fingerprint of user's second RSA public key. | 
| login_name | string | X | √ | Login name of the user. | 
| created_on | timestamp | X | √ | Timestamp when the user was created. | 
| custom_landing_page_url | string | X | √ | Snowflake Support is allowed to use the user or account. | 
| expires_at_time | timestamp | X | √ | The date and time when the user's status is set to EXPIRED and the user can no longer log in. | 
| locked_until_time | timestamp | X | √ | Specifies the number of minutes until the temporary lock on the user login is cleared. | 
| rsa_public_key_fp | string | X | √ | Fingerprint of user's RSA public key. | 
| rsa_public_key_2 | string | X | √ | Second RSA public key of the user. | 
| has_rsa_public_key | bool | X | √ | Whether the user has RSA public key. | 
| days_to_expiry | string | X | √ | User record will be treated as expired after specified number of days. | 
| ext_authn_duo | bool | X | √ | Whether Duo Security is enabled as second factor authentication. | 
| mins_to_bypass_network_policy | string | X | √ | Temporary bypass network policy on the user for a specified number of minutes. | 
| must_change_password | string | X | √ | User must change the password. | 
| email | string | X | √ | Email address of the user | 
| ext_authn_uid | string | X | √ | External authentication ID of the user. | 
| mins_to_bypass_mfa | string | X | √ | Temporary bypass MFA for the user for a specified number of minutes. | 
| snowflake_lock | string | X | √ | Whether the user or account is locked by Snowflake. | 
| owner | string | X | √ | Owner of the user in Snowflake. | 
| default_warehouse | string | X | √ | Default warehouse for this user. | 
| last_success_login | timestamp | X | √ | Date and time when the user last logged in to the Snowflake. | 
| password_last_set_time | string | X | √ | The timestamp on which the last non-null password was set for the user. Default to null if no password has been set yet. | 
| rsa_public_key | string | X | √ | RSA public key of the user. | 
| account | string | X | √ | The Snowflake account ID. | 


