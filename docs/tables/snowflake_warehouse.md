# Table: snowflake_warehouse

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| max_cluster_count | int | X | √ | Maximum number of warehouses for the (multi-cluster) warehouse (always 1 for single warehouses). | 
| running | int | X | √ | Number of SQL statements that are being executed by the warehouse. | 
| created_on | timestamp | X | √ | Date and time when the warehouse was created. | 
| resumed_on | timestamp | X | √ | Date and time when the warehouse was last started or restarted. | 
| resource_monitor | string | X | √ | ID of resource monitor explicitly assigned to the warehouse; controls the monthly credit usage for the warehouse. | 
| quiescing | string | X | √ | Percentage of the warehouse compute resources that are executing SQL statements, but will be shut down once the queries complete. | 
| type | string | X | √ | Warehouse type; STANDARD is the only currently supported type. | 
| started_clusters | int | X | √ | Number of warehouses currently started. | 
| is_default | string | X | √ | Whether the warehouse is the default for the current user. | 
| auto_suspend | int | X | √ | Specifies the number of seconds of inactivity after which a warehouse is automatically suspended. | 
| auto_resume | bool | X | √ | Specifies whether to automatically resume a warehouse when a SQL statement (e.g. query) is submitted to it. | 
| available | string | X | √ | Percentage of the warehouse compute resources that are provisioned and available. | 
| provisioning | string | X | √ | Percentage of the warehouse compute resources that are in the process of provisioning. | 
| updated_on | timestamp | X | √ | Date and time when the warehouse was last updated, which includes changing any of the properties of the warehouse or changing the state (STARTED, SUSPENDED, RESIZING) of the warehouse. | 
| owner | string | X | √ | Role that owns the warehouse. | 
| scaling_policy | string | X | √ | Policy that determines when additional warehouses (in a multi-cluster warehouse) are automatically started and shut down. | 
| min_cluster_count | int | X | √ | Minimum number of warehouses for the (multi-cluster) warehouse (always 1 for single warehouses). | 
| queued | int | X | √ | Number of SQL statements that are queued for the warehouse. | 
| other | string | X | √ | Percentage of the warehouse compute resources that are in a state other than available, provisioning, or quiescing. | 
| region | string | X | √ | The Snowflake region in which the account is located. | 
| name | string | X | √ | Name for warehouse. | 
| state | string | X | √ | Whether the warehouse is active/running (STARTED), inactive (SUSPENDED), or resizing (RESIZING). | 
| size | string | X | √ | Size of the warehouse (X-Small, Small, Medium, Large, X-Large, etc.) | 
| is_current | string | X | √ | Whether the warehouse is in use for the session. | 
| comment | string | X | √ | Comment for the warehouse. | 
| account | string | X | √ | The Snowflake account ID. | 


