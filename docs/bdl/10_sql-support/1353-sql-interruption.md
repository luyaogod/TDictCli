---
title: "SQL Interruption"
source: "fgl-topics/c_fgl_odiagmys_033.html"
breadcrumb: "SQL support > SQL database guides > Oracle® MySQL / MariaDB > BDL programming > SQL Interruption"
type: "concept"
description: "Informix® With Informix, it is possible to interrupt a long running query if the SQL INTERRUPT ON option. Oracle® MySQL and MariaDB MySQL and MariaDB provides the KILL QUERY command to interrupt a ..."
---

# SQL Interruption

## Informix®

With Informix, it is possible to interrupt a long
running query if the [SQL INTERRUPT ON](0993-using-sql-interruption.md "Interrupt long running SQL queries, or interrupt queries waiting for locked data.") option.

## Oracle® MySQL and MariaDB

MySQL and MariaDB provides the `KILL QUERY` command to interrupt a running query
on the server.

> **Note:**
>
> The database client program must open a second connection to execute the KILL QUERY
> statement.

## Solution

SQL interruption is supported with MySQL.

The database driver opens a second connection to the server and sends a `KILL
QUERY` command, with the MySQL process id of the current connection.
> **Important:**
>
> Opening a second connection does not work when using Unix sockets, connect to MySQL with a host
> name and TCP port.

## Related links

**Related concepts**  

[Using SQL interruption](0993-using-sql-interruption.md "Interrupt long running SQL queries, or interrupt queries waiting for locked data.")
