---
title: "SQL Interruption"
source: "fgl-topics/c_fgl_odiagora_050.html"
breadcrumb: "SQL support > SQL database guides > Oracle® Database > BDL programming > SQL Interruption"
type: "concept"
description: "Informix® With Informix, it is possible to interrupt a long running query if the SQL INTERRUPT ON option. ORACLE Oracle supports SQL Interruption: The db client must issue an OCIBreak() OCI call to ..."
---

# SQL Interruption

## Informix®

With Informix, it is possible to interrupt a long
running query if the [SQL INTERRUPT ON](0993-using-sql-interruption.md "Interrupt long running SQL queries, or interrupt queries waiting for locked data.") option.

## ORACLE

Oracle supports SQL Interruption: The db client must issue an `OCIBreak()` OCI
call to interrupt a query.

## Solution

The ORACLE database driver
supports SQL interruption and converts the native SQL error code
-1013 to the Informix error
code -213.

## Related links

**Related concepts**  

[Using SQL interruption](0993-using-sql-interruption.md "Interrupt long running SQL queries, or interrupt queries waiting for locked data.")
