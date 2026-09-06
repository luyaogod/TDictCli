---
title: "SQL Interruption"
source: "fgl-topics/c_fgl_odiagmsv_045.html"
breadcrumb: "SQL support > SQL database guides > Microsoft™ SQL Server > BDL programming > SQL Interruption"
type: "concept"
description: "Informix® With Informix, it is possible to interrupt a long running query if the SQL INTERRUPT ON option. Microsoft™ SQL Server Microsoft SQL Server supports SQL Interruption: The db client must issue ..."
---

# SQL Interruption

## Informix®

With Informix, it is possible to interrupt a long
running query if the [SQL INTERRUPT ON](0993-using-sql-interruption.md "Interrupt long running SQL queries, or interrupt queries waiting for locked data.") option.

## Microsoft™ SQL Server

Microsoft SQL Server supports SQL Interruption: The db
client must issue an `SQLCancel()` ODBC call to interrupt a query.

## Solution

The SQL Server database drivers support SQL interruption and return the Informix error code -213, when the statement is interrupted.

## Related links

**Related concepts**  

[Using SQL interruption](0993-using-sql-interruption.md "Interrupt long running SQL queries, or interrupt queries waiting for locked data.")
