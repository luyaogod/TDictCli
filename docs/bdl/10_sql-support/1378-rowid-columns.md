---
title: "ROWID columns"
source: "fgl-topics/c_fgl_odiagora_007.html"
breadcrumb: "SQL support > SQL database guides > Oracle® Database > Data dictionary > ROWID columns"
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

## ORACLE

Oracle® supports `ROWID`
columns, but the data type is different from Informix
ROWIDs: Oracle rowids are
`CHAR(18)`.

For example:

```
AAAA8mAALAAAAQkAAA
```

Since Oracle rowids are physical
addresses, they cannot be used as permanent row identifiers: After a `DELETE`, an
`INSERT` statement might reuse the physical place of the deleted row, to store the
new row.

The Oracle Call Level (OCI) provides functions to get the rowid of the last row related to an
`INSERT`, `UPDATE`, `DELETE` statement, ora
`FETCH` of a cursor declared with a `SELECT FOR UPDATE`.

## Solution

If your Genero BDL application uses Informix rowid
columns, review the program logic to use the primary keys when available, or use Oracle rowids.

To hold Oracle rowids, the type of variables containing `ROWID` values must be
changed to `CHAR(18)`.
> **Tip:**
>
> Informix `INTEGER` rowids fit in a
> `CHAR(18)` variable. If you adapt the code to use `CHAR(18)` rowids,
> the code will still work with Informix database. However, Informix `BIGINT` rowids
> will not fit into a `CHAR(18)`, you must use a `CHAR(20)`.

When connected to an Oracle database,
all references to `sqlca.sqlerrd[6]` must be reviewed, because this register can not
contain the rowid of the last affected row. However, after an `INSERT`,
`UPDATE` or `DELETE` statement, or after a `FETCH` of a
cursor declared with a `SELECT FOR UPDATE`, the `sqlca.sqlerrm`
register is filled with the character extended (base 64) representation of the rowid of the last
affected or fetched row.
> **Note:**
>
> The sqlca.sqlerrm register only contains a single rowid, even if
> multiple rows were affected by the `INSERT`, `UPDATE` or
> `DELETE` statement.

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
