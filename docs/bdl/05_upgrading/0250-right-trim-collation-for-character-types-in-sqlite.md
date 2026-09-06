---
title: "Right-trim collation for character types in SQLite"
source: "fgl-topics/c_fgl_Migrate_to_240_sqlite_rtrim.html"
breadcrumb: "Upgrading > Upgrade Guides for Genero BDL > BDL 2.40 upgrade guide > Right-trim collation for character types in SQLite"
type: "concept"
---

# Right-trim collation for character types in SQLite

> CHAR and VARCHAR columns in SQLite need to be defined with a TRIM collation to ignore trailing spaces in comparisons.

Since version 2.40, the SQLite database driver adds the
COLLATE RTRIM keywords after the CHAR(N) and VARCHAR(N) types in
CREATE TABLE statements, when
Informix® emulation
is enabled (the default). This collation clause forces SQLite to
use right-trim comparison rules instead of the default binary mode.
The binary mode requires to have the same number of trailing spaces
in both character values to be equal. By using COLLATE RTRIM clause,
the trailing blanks are trimmed and thus ignored. it is recommended that you also
use [VAR]CHAR(N) COLLATE RTRIM in database creation scripts.

## Related links

**Related concepts**  

[SQLite](../10_sql-support/1466-sqlite.md "SQLite")
