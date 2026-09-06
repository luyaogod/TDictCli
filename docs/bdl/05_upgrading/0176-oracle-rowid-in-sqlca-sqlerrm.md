---
title: "ORACLE rowid in sqlca.sqlerrm"
source: "fgl-topics/c_fgl_Migrate_to_320_ora_rowid_sqlerrm.html"
breadcrumb: "Upgrading > Upgrade Guides for Genero BDL > BDL 3.20 upgrade guide > ORACLE rowid in sqlca.sqlerrm"
type: "concept"
---

# ORACLE rowid in sqlca.sqlerrm

> With ORACLE, the rowid of the last affected row is available in sqlca.sqlerrm.

Before Genero 3.20, the character extended (base 64) representation of the rowid of the row
affected by an `INSERT`, `UPDATE` or `DELETE` could be
found in the `sqlca.sqlerrm` register, when the following FGLPROFILE entry was
defined:

```
dbi.database.mydbname.ora.rowid.retrieve = true
```

This FGLPROFILE entry was required to fill the `sqlca.sqlerrm` register with the
result of a `SELECT ROWIDTOCHAR(?) FROM DUAL` statement, that was executed after each
`INSERT`, `UPDATE` or `DELETE` SQL statement. For
performance reasons, the `database.dbname.ora.rowid.retrieve`
switch was by default set to `false`.

> **Note:**
>
> For maximum SQL portability, replace rowids usage by primary keys. See [Using ROWID columns](../10_sql-support/1042-using-rowid-columns.md "Automatic ROWID columns is not a common database feature.").

Starting with Genero 3.20, to fill the `sqlca.sqlerrm` register, the Oracle ODI
driver uses now the `OCIRowidToChar()` API to convert the binary rowid value to the
character extended (base 64) representation. Since the additional `SELECT` statement
is not required anymore, the FGLPROFILE entry
`database.dbname.ora.rowid.retrieve` does no longer exist.

The `sqlca.sqlerrm` is now implicitly filled with the rowid of the row affected
by an `INSERT`, `UPDATE`, `DELETE` statement, and after
a `FETCH` of a cursor declared with `SELECT FOR UPDATE`.

## Related links

**Related concepts**  

[ROWID columns](../10_sql-support/1378-rowid-columns.md "ROWID columns")
