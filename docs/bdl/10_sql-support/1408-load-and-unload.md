---
title: "LOAD and UNLOAD"
source: "fgl-topics/c_fgl_odiagora_043.html"
breadcrumb: "SQL support > SQL database guides > Oracle® Database > BDL programming > LOAD and UNLOAD"
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

## ORACLE

ORACLE provides tools like SQL\*Plus and SQL\*Loader to load/unload data from a database.

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

The `LOAD` and `UNLOAD` BDL instructions are supported with ORACLE
with some limitations:

- There is a difference when using ORACLE `DATE` columns. `DATE`
  columns created in the ORACLE database are equivalent to Informix `DATETIME YEAR TO SECOND` columns. In `LOAD` and
  `UNLOAD`, all ORACLE `DATE` columns are treated as Informix `DATETIME YEAR TO SECOND` columns and thus will be
  unloaded with the "`YYYY-MM-DD hh:mm:ss`" format.
- Informix `INTEGER` and
  `SMALLINT` are mapped to ORACLE `NUMBER(?)` columns. Those values will
  be unloaded as Informix `DECIMAL(10)` and
  `DECIMAL(5)` values, that is, with a trailing dot-zero ".0".
- When using an Informix database, simple dates are
  unloaded using the DBDATE format (ex:"23/12/1998"). Unloading from an Informix database for loading into an ORACLE database is not
  supported.

## Related links

**Related concepts**  

[LOAD and UNLOAD instructions](1038-load-and-unload-instructions.md "The LOAD and UNLOAD instructions can produce different data formats depending on the database server type.")
