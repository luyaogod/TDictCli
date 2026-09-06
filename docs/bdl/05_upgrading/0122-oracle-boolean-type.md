---
title: "ORACLE BOOLEAN type"
source: "fgl-topics/c_fgl_Migrate_to_500_oracle_boolean.html"
breadcrumb: "Upgrading > Upgrade Guides for Genero BDL > BDL 5.00 upgrade guide > ORACLE® BOOLEAN type"
type: "concept"
---

# ORACLE BOOLEAN type

> Support for native Oracle SQL BOOLEAN type.

Starting with Oracle version 23c, the Oracle SQL language supports a native
`BOOLEAN` data type in DDL amd DML statements. Prior to 23c, the
`BOOLEAN` type was only available in PL/SQL.

Genero BDL 5.00 introduces a new ODI driver dbmora\_23 for Oracle 23c, that supports the native
SQL `BOOLEAN` data type to store FGL `BOOLEAN` values. This is the new
recommended SQL type mapping.

For backward compatibility, if your database defines already `CHAR(1)` columns to
store FGL `BOOLEAN` values, set the following FGLPROFILE entry, to get the former
`BOOLEAN/CHAR(1)`
conversion:

```
dbi.database.stores.ora.boolean.aschar = true
```

> **Important:**
>
> The `BOOLEAN/CHAR(1)` conversion rule applies to all SQL tables for a given SQL
> connection. If you have existing boolean data in several tables using `CHAR(1)`
> columns, you need to plan for an upgrade process in order to use the new native SQL
> `BOOLEAN` type of Oracle.

For more details, see [BOOLEAN data type](../10_sql-support/1372-boolean-data-type.md).
