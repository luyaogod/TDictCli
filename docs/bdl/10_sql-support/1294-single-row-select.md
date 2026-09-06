---
title: "Single row SELECT"
source: "fgl-topics/c_fgl_odiagmsv_028.html"
breadcrumb: "SQL support > SQL database guides > Microsoft™ SQL Server > Data manipulation > Single row SELECT"
type: "concept"
description: "Informix® With Informix, you must use the system table with a condition on the table id: SELECT USER FROM systables WHERE tabid=1 Microsoft™ SQL Server With SQL Server, omit the FROM clause to ..."
---

# Single row SELECT

## Informix®

With Informix, you must use the system table with a
condition on the table id:

```
SELECT USER FROM systables WHERE tabid=1
```

## Microsoft™ SQL Server

With SQL Server, omit the `FROM` clause to generate one row
only:

```
SELECT CURRENT_USER
```

## Solution

Check the BDL sources for "`FROM systables WHERE tabid=1`" and use dynamic SQL to
resolve this problem.

Consider writing a `FUNCTION` which produces the `FROM` and
`WHERE` part, depending on the target database type.
