---
title: "ROWID columns"
source: "fgl-topics/c_fgl_odiagpgs_007.html"
breadcrumb: "SQL support > SQL database guides > PostgreSQL > Data dictionary > ROWID columns"
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

## PostgreSQL

Before PostgreSQL 12, OID columns (similar to Informix ROWID columns) were available. Starting
with PostgreSQL version 12, `OID` columns are no longer supported.

PostgreSQL tables have a CTID system column, but the value of this column defines the storage
location of the row, the value changes when updating.

## Solution

Starting with PostgreSQL 12, user-table OIDs are no longer supported. Consequently, Informix
ROWID columns cannot be emulated.

As a general SQL programming pattern, get rid of ROWID usage and use `PRIMARY KEY`
columns to identify database table rows.

## Related links

**Related concepts**  

[Using ROWID columns](1042-using-rowid-columns.md "Automatic ROWID columns is not a common database feature.")
