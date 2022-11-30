# Table: snowflake_database

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| account | string | X | √ | The Snowflake account ID. | 
| name | string | X | √ | Name of the database. | 
| created_on | timestamp | X | √ | Creation time of the database. | 
| is_default | string | X | √ | Name of the default database for authenticating user. | 
| origin | string | X | √ | Name of the origin database. | 
| region | string | X | √ | The Snowflake region in which the account is located. | 
| comment | string | X | √ | Comment for this database. | 
| is_current | string | X | √ | Name of the current database for authenticating user. | 
| options | string | X | √ |  | 
| owner | string | X | √ | Name of the role that owns the schema. | 
| retention_time | int | X | √ | Number of days that historical data is retained for Time Travel. | 


