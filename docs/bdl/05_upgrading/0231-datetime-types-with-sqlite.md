---
title: "DATETIME types with SQLite"
source: "fgl-topics/c_fgl_Migrate_to_251_sqlite_datetime.html"
breadcrumb: "Upgrading > Upgrade Guides for Genero BDL > BDL 2.51 upgrade guide > DATETIME types with SQLite"
type: "concept"
description: "Better support for Informix DATETIME types emulation within SQLite."
---

# DATETIME types with SQLite

> Better support for Informix® DATETIME types emulation within SQLite.

Before version 2.51, `DATETIME` SQL types where converted to SQLite types as
follows:

- `DATETIME HOUR TO SECOND` type was translated to TIME (hh:mm:ss).
- `DATETIME YEAR TO FRACTION` and all other combinations (except `HOUR TO
  SECOND`) were translated to `TIMESTAMP` (YYYY-MM-DD hh:mm:ss.fff).

Since most `DATETIME` types were converted to `TIMESTAMP`, it was
not possible to distinguish common date/time types such as `DATETIME HOUR TO MINUTE`
or `DATETIME YEAR TO MINUTE`, especially when extracting the database schema with
fgldbsch. Type information was lost and this prevented schema-base variable
definitions with `DEFINE LIKE`.

Starting with version 2.51, common `DATETIME` SQL types are now mapped to
different types in SQLite, to provide better support for these types. In fact, SQLite allows you to
define table columns with custom types (you can use any type name). However, the number of tokens in
the syntax is limited so it's not possible to use, for example, the tokens `DATETIME YEAR TO
SECOND` directly. The Genero database driver uses this SQLite SQL language feature to map
Informix-style DATETIME types to specific custom types. For example, a `DATETIME HOUR TO
MINUTE` becomes a `SMALLTIME`, a `DATETIME YEAR TO
FRACTION(2)` becomes a `DATETIME(2)`, etc. Furthermore, the data values
inserted in the database now match exactly the precision of the original `DATETIME`
type. For more details about date/time mapping and emulation, see [DATE and DATETIME data types](../10_sql-support/1481-date-and-datetime-data-types.md).

## Related links

**Related concepts**  

[SQLite](../10_sql-support/1466-sqlite.md "SQLite")

[fgldbsch](../13_programming-tools/2521-fgldbsch.md "The fgldbsch tool generates the database schema files from an existing database.")
