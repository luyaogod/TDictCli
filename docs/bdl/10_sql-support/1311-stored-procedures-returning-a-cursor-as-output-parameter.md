---
title: "Stored procedures returning a cursor as output parameter"
source: "fgl-topics/c_fgl_sql_programming_028.html"
breadcrumb: "SQL support > SQL database guides > Microsoft™ SQL Server > BDL programming > Stored procedure calls > Stored procedures returning a cursor as output parameter"
type: "concept"
description: "SQL Server supports \"cursor output parameters\": A stored procedure can declare/open a cursor and return a reference of the cursor to the caller. SQL Server stored procedures returning a cursor as ..."
---

# Stored procedures returning a cursor as output parameter

SQL Server supports "cursor output parameters": A stored procedure can declare/open a cursor and
return a reference of the cursor to the caller.

SQL Server stored procedures returning a cursor as output parameter are not supported. There are
two reasons for this: The language does not have a data type to store a server cursor
reference, and the underlying ODBC driver does not support this anyway.
