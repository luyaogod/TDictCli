---
title: "Static SQL cache is removed"
source: "fgl-topics/c_fgl_Migrate_to_200_static_sql_cache.html"
breadcrumb: "Upgrading > Upgrade Guides for Genero BDL > BDL 2.00 upgrade guide > Static SQL cache is removed"
type: "concept"
---

# Static SQL cache is removed

> The Static SQL Cache has been removed.

Before version 2.00, the size of the static SQL cache is defined by an [FGLPROFILE](../07_configuration/0486-fglprofile-entries-for-core-language.md "This is a summary of FGLPROFILE entries supported by the core BDL.") entry:

```
dbi.sql.static.optimization.cache.size = max
```

This entry was provided to optimize SQL execution without touching code using a lot of static SQL
statements, especially when using non-Informix® database servers where the preparation of static SQL statements is slower than with
Informix. This is useful for fast migrations, but there
were a lot of side effects and unexpected errors. On the other hand, recent database client
libraries provide a "deferred prepare" option, to postpone the preparation of the SQL statement to
the first statement execution, making prepared SQL cache useless.

Starting with version 2.00, the Static SQL Cache has been removed for the reasons described.
Programs continue to run without changing the code. With some database brands, SQL statement
execution can be optimized by using `PREPARE` / `EXECUTE`, when the
SQL statement needs to be executed many times. In such case, best performance improvement can be
achieved by using a transaction block (`BEGIN WORK` / `COMMIT WORK`),
and thus avoid auto-commits at each `EXECUTE`.

## Related links

**Related concepts**  

[Dynamic SQL management](../10_sql-support/1140-dynamic-sql-management.md "Explains how to execute and manage SQL statements at runtime.")
