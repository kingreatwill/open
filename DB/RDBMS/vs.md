

### 慢日志

mssql

```sql
SELECT creation_time  N'语句编译时间'
        ,last_execution_time  N'上次执行时间'
        ,total_physical_reads N'物理读取总次数'
        ,total_logical_reads/execution_count N'每次逻辑读次数'
        ,total_logical_reads  N'逻辑读取总次数'
        ,total_logical_writes N'逻辑写入总次数'
        ,execution_count  N'执行次数'
        ,total_worker_time/1000 N'所用的CPU总时间ms'
        ,total_elapsed_time/1000  N'总花费时间ms'
        ,(total_elapsed_time / execution_count)/1000  N'平均时间ms'
        ,SUBSTRING(st.text, (qs.statement_start_offset/2) + 1,
         ((CASE statement_end_offset
          WHEN -1 THEN DATALENGTH(st.text)
          ELSE qs.statement_end_offset END
            - qs.statement_start_offset)/2) + 1) N'执行语句'
FROM sys.dm_exec_query_stats AS qs
CROSS APPLY sys.dm_exec_sql_text(qs.sql_handle) st
where SUBSTRING(st.text, (qs.statement_start_offset/2) + 1,
         ((CASE statement_end_offset
          WHEN -1 THEN DATALENGTH(st.text)
          ELSE qs.statement_end_offset END
            - qs.statement_start_offset)/2) + 1) not like '%fetch%'
ORDER BY  total_elapsed_time / execution_count DESC;
```
> 需要使用master数据库,use master;

mysql
`show variables like '%query%';`
参数名称	|状态	|说明
|---|---|---|
long_query_time | 1 | 设置记录查询语句时间的阈值, 单位 (S)
slow_query_log | ON| 慢查询日志的开关
slow_query_log_file | /usr/local/var/mysql/slow.log | 慢查询日志的路径,文件路径可以不用设置,如果mysql对当前路径不能写,日志开关不起作用

PostgreSQL
`show log_min_duration_statement;`

参数名称	|状态	|说明
|---|---|---|
logging_collector	|on	|Start a subprocess to capture stderr output and/or csvlogs into log files.
log_min_duration_statement	|1s	|Sets the minimum execution time above which statements will be logged.
log_filename	|postgresql-%Y-%m-%d_%H%M%S.log	|Sets the file name pattern for log files.