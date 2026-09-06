---
title: "IBM Informix SQL ANSI Mode"
source: "fgl-topics/c_fgl_sql_programming_077.html"
breadcrumb: "SQL support > SQL programming > SQL portability > IBM® Informix® SQL ANSI Mode"
type: "concept"
description: "Understanding the impact of the SQL ANSI mode of IBM Informix."
---

# IBM Informix SQL ANSI Mode

> Understanding the impact of the SQL ANSI mode of IBM® Informix®.

IBM Informix
allows you to create databases in ANSI mode, supposed to follow the ANSI SQL specification.

If you are not using the ANSI mode with IBM Informix, we suggest you keep the database as is, because turning an IBM Informix database
into ANSI mode can result in unexpected behavior of the programs.

Here are some ANSI mode issues extracted from the IBM
Informix books:

- Some actions, like `CREATE INDEX` will generate a warning but will not be
  forbidden.
- Buffered logging is not allowed to enforce data recovery. (Buffered logging provides better
  performance)
- The table-naming scheme allows different users to create tables without having to worry about
  name conflicts.
- Owner is required in DB object names (`SELECT ... FROM "owner".table` ). You
  must quote the owner name to prevent automatic translation of the owner name into uppercase:
  `SELECT ... FROM owner.table` becomes `SELECT .. FROM OWNER.table`
  and thus, the table is not found in the database.
- Default privileges differ: When creating a table, the server grants privileges to the table
  owner and the DBA only. The same thing happens for the 'Execute' privilege when creating stored
  procedures.
- Default isolation level is `REPEATABLE READ`.
- An error is generated if any character field is filled with a value that is longer than the
  field width.
- `DECIMAL(p)` (floating point decimals) are automatically converted to
  `DECIMAL(p,0)` (fixed point decimals).
- Closing a closed cursor generates an SQL error.

It will take more time to adapt the programs to the IBM
Informix ANSI mode than using the database interface
to simulate the native mode of IBM Informix.
