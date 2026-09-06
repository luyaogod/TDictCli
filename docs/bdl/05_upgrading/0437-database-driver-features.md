---
title: "Database driver features"
source: "fgl-topics/c_fgl_Mig0000_035.html"
breadcrumb: "Upgrading > Migrating from Four Js BDS to Genero BDL > 4GL Programming topics > Database driver features"
type: "concept"
---

# Database driver features

> Some ODI features are no longer supported in Genero BDL.

The following ODI features are no longer supported in Genero:

- The concept of "ODI connectors" (as sub-processes) has been replaced by "ODI drivers", which are
  shared libraries loaded by the runtime system. See [Database driver specification (driver)](../10_sql-support/1073-database-driver-specification-driver.md) for
  more details about ODI drivers.
- A cache of prepared SQL statements could be used to improve Static SQL execute. For more
  details, see [Static SQL cache desupport in
  Genero 2.00](0319-static-sql-cache-is-removed.md "The Static SQL Cache has been removed.").
- SQL Directives could be defined for each database brand, to use a specific SQL syntax. For more
  details, see [SQL directive sets desupport
  in Genero 2.00](0320-sql-directive-set-removed.md "The SQL directive set specification has been removed.")
