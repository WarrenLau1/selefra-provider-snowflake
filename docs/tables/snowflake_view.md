# Table: snowflake_view

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| is_materialized | bool | X | √ | True if the view is a materialized view; false otherwise. | 
| name | string | X | √ | The name of the view. | 
| database_name | string | X | √ | The name of the database in which the view exists. | 
| created_on | timestamp | X | √ | The timestamp at which the view was created. | 
| owner | string | X | √ | The owner of the view. | 
| text | string | X | √ | The text of the command that created the view, e.g., CREATE VIEW. | 
| schema_name | string | X | √ | The name of the schema in which the view exists. | 
| comment | string | X | √ | Optional comment. | 
| is_secure | bool | X | √ | True if the view is a secure view; false otherwise. | 
| region | string | X | √ | The Snowflake region in which the account is located. | 
| account | string | X | √ | The Snowflake account ID. | 


