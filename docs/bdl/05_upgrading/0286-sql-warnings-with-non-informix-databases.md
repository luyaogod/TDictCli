---
title: "SQL Warnings with non-Informix databases"
source: "fgl-topics/c_fgl_Migrate_to_220_sql_warnings.html"
breadcrumb: "Upgrading > Upgrade Guides for Genero BDL > BDL 2.20 upgrade guide > SQL Warnings with non-Informix databases"
type: "concept"
---

# SQL Warnings with non-Informix databases

> SQL Warnings are now propagated for all database drivers, and can set the sqlca.sqlawarn, SQLSTATE and SQLERRMESSAGE registers.

Before version 2.20, is was impossible for a non-Informix driver to return SQL Warning
information in sqlca, SQLSTATE and SQLERRMESSAGE.

This new behavior will have no impact if you test SQL Errors with status or sqlca.sqlcode, as
these registers remain zero if an SQL Warning is raised. However, if you are using SQLSTATE
to check for SQL Errors, you must now distinguish SQLSTATE of class 01: These are SQL
Warnings, not SQL errors.

For example, with some database engines, a `DELETE` statement that removes all
rows from a table will set SQLSTATE to `"01504"`. The test `SQLSTATE <>
"00000"` will be false, and run into the error handling block, which is
unexpected:

```
MAIN
   DATABASE stores
   WHENEVER ERROR CONTINUE
   DELETE FROM customer 
   IF SQLSTATE <> "00000" THEN
      -- handle error 
   END IF
END MAIN
```

To check for successful SQL execution with or without warning,
you can, for example, code:

```
MAIN
   DATABASE stores 
   WHENEVER ERROR CONTINUE
   DELETE FROM customer 
   IF NOT (SQLSTATE=="00000" OR SQLSTATE MATCHES "01*") THEN
      -- handle error 
   END IF
END MAIN
```
