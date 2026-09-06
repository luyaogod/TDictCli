---
title: "ROWID columns"
source: "fgl-topics/c_fgl_odiagmys_007.html"
breadcrumb: "SQL support > SQL database guides > Oracle® MySQL / MariaDB > Data dictionary > ROWID columns"
type: "concept"
description: "Informix® When creating a table, Informix automatically adds a ROWID integer column (applies to non-fragmented tables only). The ROWID column is auto-filled with a unique number and can be used like a ..."
---

# ROWID columns

## Informix®

When creating a table, Informix automatically adds a
`ROWID` integer column (applies to non-fragmented tables only).

The `ROWID` column is auto-filled with a unique number and can be used like a
primary key to access a given row.

Starting with Informix version 15, when using the "large tables" option, rowid columns are
defined as 8 byte `BIGINT` integers. When using "small tables" option (onconfig
`TABLE_SIZE SMALL`), the rowids are defined as 4 byte `INTEGER`.

> **Note:**
>
> Informix `ROWID` usage was a
> common practice in the early days of Informix
> 4GL programming. Today it is recommended to define all your database tables with a
> `PRIMARY KEY` to uniquely identify rows.

With Informix, the `sqlca.sqlerrd[6]`
register contains the `ROWID` of the last row affected by an `INSERT`,
`UPDATE` or `DELETE` statement.

## Oracle® MySQL and MariaDB

MySQL and MariadDB do not support implicit rowid columns like Informix does.

Starting with version 8.0.30, MySQL provides Generated Invisible Primary Keys
(GIPKs): When setting the `sql_generate_invisible_primary_key` server variable
to `ON`, new created tables will implicitly get a `my_row_id` column
as an `AUTO_INCREMENT INVISIBLE PRIMARY KEY`. This could be used to emulate Informix
rowids, but is not supported by Genero, as it would require to adapt the [SERIAL emulation](1331-serial-and-bigserial-data-type.md): There can only be one auto-incremented
column, and one primary key in a MySQL database table.

## Solution

If your Genero BDL
application uses rowid columns, review the program logic to use primary keys insead. If the database
table does no define a primary key, it should be added. All references to
`SQLCA.SQLERRD[6]` must be removed, because this variable will not hold the
`ROWID` of the last modified row.

For databases where the keyword of the
rowid pseudo-column is different than "`ROWID`", the translation can be controlled
with the following FGLPROFILE
entry:

```
dbi.database.dsname.ifxemul.rowid = { true | false }
```

For more details see [IBM Informix emulation parameters in FGLPROFILE](1079-ibm-informix-emulation-parameters-in-fglprofile.md "Emulation of Informix specific SQL features can be controlled with FGLPROFILE entries.").

## Related links

**Related concepts**  

[Using ROWID columns](1042-using-rowid-columns.md "Automatic ROWID columns is not a common database feature.")
