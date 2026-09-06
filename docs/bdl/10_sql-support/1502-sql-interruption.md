---
title: "SQL Interruption"
source: "fgl-topics/c_fgl_odiagsqt_043.html"
breadcrumb: "SQL support > SQL database guides > SQLite > BDL programming > SQL Interruption"
type: "concept"
description: "Informix® With Informix, it is possible to interrupt a long running query if the SQL INTERRUPT ON option. SQLite SQLite supports SQL Interruption: The db client must issue an sqlite3_interrupt() ODBC ..."
---

# SQL Interruption

## Informix®

With Informix, it is possible to interrupt a long
running query if the [SQL INTERRUPT ON](0993-using-sql-interruption.md "Interrupt long running SQL queries, or interrupt queries waiting for locked data.") option.

## SQLite

SQLite supports SQL Interruption: The db client must issue an
`sqlite3_interrupt()` ODBC call to interrupt a query.

## Solution

The SQLite database driver supports SQL interruption and converts the native SQL execution status
`SQLITE_ABORT` to the Informix error code
-213.

## Related links

**Related concepts**  

[Using SQL interruption](0993-using-sql-interruption.md "Interrupt long running SQL queries, or interrupt queries waiting for locked data.")
