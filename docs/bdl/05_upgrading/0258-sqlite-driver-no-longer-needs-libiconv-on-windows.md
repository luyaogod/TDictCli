---
title: "SQLite driver no longer needs libiconv on Windows"
source: "fgl-topics/c_fgl_Migrate_to_232_sqlite_libiconv.html"
breadcrumb: "Upgrading > Upgrade Guides for Genero BDL > BDL 2.32 upgrade guide > SQLite driver no longer needs libiconv on Windows®"
type: "concept"
---

# SQLite driver no longer needs libiconv on Windows

> UTF-8 string data storage in SQLite requires conversion when the application is not UTF-8.

Starting with version 2.32, the SQLite driver (dbmsqt3xx) no longer needs the
LIBICONV.DLL library on Windows® platforms to do charset conversion, when the application locale is not
UTF-8.

## Related links

**Related concepts**  

[SQLite](../10_sql-support/1466-sqlite.md "SQLite")

[Localization](../09_advanced-features/0863-localization.md "Localization support allows you to implement programs that follow specific language and cultural rules.")
