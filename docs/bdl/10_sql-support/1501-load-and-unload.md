---
title: "LOAD and UNLOAD"
source: "fgl-topics/c_fgl_odiagsqt_025.html"
breadcrumb: "SQL support > SQL database guides > SQLite > BDL programming > LOAD and UNLOAD"
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

## SQLite

SQLite does not natively provide `LOAD` and `UNLOAD`
instructions.

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

The `LOAD` and `UNLOAD` BDL instructions are supported with
SQLite.

## Related links

**Related concepts**  

[LOAD and UNLOAD instructions](1038-load-and-unload-instructions.md "The LOAD and UNLOAD instructions can produce different data formats depending on the database server type.")
