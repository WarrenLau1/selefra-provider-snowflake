# Table: snowflake_resource_monitor

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| end_time | timestamp | X | √ | Date and time when the monitor was stopped. | 
| suspend_immediately_at | json | X | √ | Levels to which to suspend warehouse. | 
| suspend_at | json | X | √ | Levels to which to suspend warehouse. | 
| owner | string | X | √ | Role that owns the warehouse. | 
| name | string | X | √ | Name for warehouse. | 
| credit_quota | float | X | √ | Specifies the number of Snowflake credits allocated to the monitor for the specified frequency interval. | 
| used_credits | float | X | √ | Number of credits used in the current monthly billing cycle by all the warehouses associated with the resource monitor. | 
| frequency | string | X | √ | The interval at which the used credits reset relative to the specified start date (Daily,Weekly,Monthly,Yearly,Never). | 
| start_time | timestamp | X | √ | Date and time when the monitor was started. | 
| notify_at | json | X | √ | Levels to which to alert. | 
| notify_users | string | X | √ | Who to notify when alerting. | 
| region | string | X | √ | The Snowflake region in which the account is located. | 
| remaining_credits | float | X | √ | Number of credits still available to use in the current monthly billing cycle. | 
| level | string | X | √ | Specifies whether the resource monitor is used to monitor the credit usage for your entire Account (i.e. all warehouses in the account) or a specific set of individual warehouses. | 
| created_on | timestamp | X | √ | Date and time when the monitor was created. | 
| comment | string | X | √ | Comment for the warehouse. | 
| account | string | X | √ | The Snowflake account ID. | 


