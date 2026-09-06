---
title: "Modifying many rows in a table"
source: "fgl-topics/c_fgl_odiagsqt_tx_perf.html"
breadcrumb: "SQL support > SQL database guides > SQLite > BDL programming > Modifying many rows in a table"
type: "concept"
description: "SQLite SQLite is very slow when doing commits, because of the technique used to ensure data integrity (see SQLite documentation for details). When a program executes a DML statement like INSERT , it ..."
---

# Modifying many rows in a table

## SQLite

SQLite is very slow when doing commits, because of the technique used to ensure data integrity
(see SQLite documentation for details).

When a program executes a DML statement like `INSERT`, it will be automatically
committed by SQLite. As result, if you do not enclose the SQL instruction between `BEGIN
WORK` / `COMMIT WORK`, there will be as many commits as data manipulation
statements.

## Solution

If a program must modify many rows in a table, execute the SQL statement within a transaction
block delimited by `BEGIN WORK` / `COMMIT WORK` instructions. This
will dramatically speed up the program with SQLite, and even with other non-Informix database
servers.

See [Performance with transactions](1054-performance-with-transactions.md "Commit database changes by blocks of transaction speeds performance with some database servers.").
