---
title: "SQL Interruption"
source: "fgl-topics/c_fgl_odiagdmg_043.html"
breadcrumb: "SQL support > SQL database guides > Dameng® database server > BDL programming > SQL Interruption"
type: "concept"
description: "Informix® With Informix, it is possible to interrupt a long running query if the SQL INTERRUPT ON option. Dameng® Dameng supports SQL Interruption: The db client must issue an dpi_cancel() ODBC call ..."
---

# SQL Interruption

## Informix®

With Informix, it is possible to interrupt a long
running query if the [SQL INTERRUPT ON](0993-using-sql-interruption.md "Interrupt long running SQL queries, or interrupt queries waiting for locked data.") option.

## Dameng®

Dameng supports SQL Interruption: The db client must
issue an `dpi_cancel()` ODBC call to interrupt a query.

## Solution

The Dameng database driver supports SQL interruption and
converts the native SQL error code -952 to the Informix
error code -213.

## Related links

**Related concepts**  

[Using SQL interruption](0993-using-sql-interruption.md "Interrupt long running SQL queries, or interrupt queries waiting for locked data.")
