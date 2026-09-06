---
title: "Inconsistent USING clauses"
source: "fgl-topics/c_fgl_Migrate_to_200_sql_using.html"
breadcrumb: "Upgrading > Upgrade Guides for Genero BDL > BDL 2.00 upgrade guide > Inconsistent USING clauses"
type: "concept"
---

# Inconsistent USING clauses

> Having data types changing at each execute is no longer supported.

> **Important:**
>
> **This issue applies to non-Informix databases
> only.**

Before version 2.00, it was possible to execute a prepared statement with the variable list
changing at each [EXECUTE](../10_sql-support/1143-execute-sql-statement.md "This instruction runs an SQL statement previously prepared.")
statement:

```
DEFINE var1 DECIMAL(6,2)
DEFINE var2 CHAR(10)
DEFINE var3 DATE
PREPARE st1 FROM "INSERT INTO tab1 VALUES ( ?. ?, ? )"
EXECUTE st1 USING var1, var2, var3
EXECUTE st1 USING var2, var3, var1   -- different order = different data types
```

The database interface of version 2.00 has been rewritten for better performance. Having data
types changing at each execute is no longer supported.

Error [-254](../15_library-reference/4483-genero-bdl-errors.md)
will be raised if different data types are used in subsequent EXECUTE
statements (with the same statement name).

## Related links

**Related concepts**  

[Database transactions](../10_sql-support/0992-database-transactions.md "Database transactions define a set of SQL instructions to be executed as a whole, or rolled back as a whole.")
