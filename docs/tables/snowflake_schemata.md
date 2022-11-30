# Table: snowflake_schemata

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| region | string | X | √ | The Snowflake region in which the account is located. | 
| is_transient | string | X | √ | Whether this is a transient schema. | 
| comment | string | X | √ | Comment for this schema. | 
| created | timestamp | X | √ | Creation time of the schema. | 
| schema_id | string | X | √ | ID of the schema. | 
| schema_owner | string | X | √ | Name of the role that owns the schema. | 
| deleted | timestamp | X | √ | Deletion time of the schema. | 
| schema_name | string | X | √ | Name of the schema. | 
| catalog_name | string | X | √ | Database that the schema belongs to. | 
| retention_time | int | X | √ | Number of days that historical data is retained for Time Travel. | 
| is_managed_access | string | X | √ | Whether the schema is a managed access schema. | 
| last_altered | timestamp | X | √ | Last altered time of the schema. | 
| account | string | X | √ | The Snowflake account ID. | 
| catalog_id | string | X | √ | ID of the database that the schema belongs to. | 


