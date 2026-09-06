---
title: "LOAD and UNLOAD"
source: "fgl-topics/c_fgl_odiagmsv_040.html"
breadcrumb: "SQL support > SQL database guides > Microsoft™ SQL Server > BDL programming > LOAD and UNLOAD"
type: "concept"
description: "Informix® Informix provides two SQL instructions to export / import data from / into a database table: The UNLOAD instruction copies rows from a database table into a text file: UNLOAD TO ..."
---

# LOAD and UNLOAD

## Informix®

Informix provides two SQL instructions to export /
import data from / into a database table:

The `UNLOAD` instruction copies rows from a database table into a text
file:

```
UNLOAD TO "filename.unl" SELECT * FROM tab1 WHERE ..
```

The
`LOAD` instructions insert rows from a text file into a database
table:

```
LOAD FROM "filename.unl" INSERT INTO tab1
```

## Microsoft™ SQL Server

Microsoft SQL Server has `LOAD` and
`UNLOAD` instructions, but those commands are related to database backup and
recovery. Do not confuse with Informix commands.

## Solution

`LOAD` and `UNLOAD` instruction are implemented in the Genero BDL
runtime system with basic `INSERT` (for `LOAD`) or `SELECT`
(for `UNLOAD`) SQL commands. The `LOAD` and `UNLOAD`
instruction can be supported with various database servers.

However, `LOAD` and `UNLOAD` require the description of the column
types in order to work, that can lead to some differences in the data formatting.

> **Note:**
>
> If no transaction is started, the `LOAD` instruction will automatically execute
> a `BEGIN WORK` and `COMMIT WORK` when finished, or `ROLLBACK
> WORK` if a row insertion failed while loading. Terminating a transaction will automatically
> close cursors not defined `WITH HOLD` option. To workaround this situation, see more
> details in the [LOAD](1176-load.md "Inserts data from a file into an existing database table.") reference topic.

The `LOAD` and `UNLOAD` BDL instructions are supported with Microsoft SQL Server with some limitations:

- The `LOAD` instruction does not work with tables using emulated
  `SERIAL` columns because the generated `INSERT` statement holds the
  `SERIAL` column which is actually a `IDENTITY` column in SQL Server.
  See the [limitations of INSERT statements when using SERIAL
  types](1277-serial-and-bigserial-data-types.md).
- Starting with Microsoft SQL Server 2008,
  Informix `DATETIME` data is
  stored in SQL Server `DATETIME2(n<=5)` or
  `TIME(n<=5)` columns, depending on the precision of the original
  `DATETIME` type:
  - With `DATETIME2(n<=5)` columns, the result of `LOAD` and
    `UNLOAD` is equivalent to Informix
    `DATETIME YEAR TO SECOND` or `DATETIME YEAR TO FRACTION(n)` columns.
    The data format will be "`YYYY-MM-DD
    hh:mm:ss[.fff...``]`", where fff...
    depends on the precision (n) of the `DATETIME2(n)` column.
  - With `TIME(n)` columns, the result of `LOAD` and
    `UNLOAD` is equivalent to Informix
    `DATETIME HOUR TO SECOND` or `DATETIME HOUR TO FRACTION(n)` columns.
    The data format will be "hh:mm:ss`[`.fff...`]`", where
    fff... depends on the precision (n) of the `TIME(n)` column.
- With an Informix database, simple dates are
  unloaded with the DBDATE format (ex:"23/12/1998"). Therefore, unloading from an Informix database for loading into a
  Microsoft SQL Server database is not supported.

## Related links

**Related concepts**  

[LOAD and UNLOAD instructions](1038-load-and-unload-instructions.md "The LOAD and UNLOAD instructions can produce different data formats depending on the database server type.")
