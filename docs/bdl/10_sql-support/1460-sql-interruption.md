---
title: "SQL Interruption"
source: "fgl-topics/c_fgl_odiagpgs_041.html"
breadcrumb: "SQL support > SQL database guides > PostgreSQL > BDL programming > SQL Interruption"
type: "concept"
description: "Informix® With Informix, it is possible to interrupt a long running query if the SQL INTERRUPT ON option. PostgreSQL PostgreSQL supports SQL Interruption: The db client must issue an PQcancel() libPQ ..."
---

# SQL Interruption

## Informix®

With Informix, it is possible to interrupt a long
running query if the [SQL INTERRUPT ON](0993-using-sql-interruption.md "Interrupt long running SQL queries, or interrupt queries waiting for locked data.") option.

## PostgreSQL

PostgreSQL supports SQL Interruption: The db client must issue an `PQcancel()`
libPQ call to interrupt a query.

## Solution

The PostgreSQL database driver supports SQL interruption and converts
the SQLSTATE code 57014 to the Informix
error code -213.

## Related links

**Related concepts**  

[Using SQL interruption](0993-using-sql-interruption.md "Interrupt long running SQL queries, or interrupt queries waiting for locked data.")
