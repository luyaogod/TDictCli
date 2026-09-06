---
title: "SQL directive set removed"
source: "fgl-topics/c_fgl_Migrate_to_200_sql_directive_set.html"
breadcrumb: "Upgrading > Upgrade Guides for Genero BDL > BDL 2.00 upgrade guide > SQL directive set removed"
type: "concept"
---

# SQL directive set removed

> The SQL directive set specification has been removed.

Before version 2.00, it was possible to define SQL directive sets in [FGLPROFILE](../07_configuration/0486-fglprofile-entries-for-core-language.md "This is a summary of FGLPROFILE entries supported by the core BDL."):

```
dbi.sqldirset.set-name.directive.directive-name.{deflist|substrg} = "value"
```

With this feature, one could write SQL statement with database specific syntax, to be adapted at
runtime depending on the target database
type:

```
-- substrg directive:
dbi.sqldirset.ora.directive.trunc.substrg = "TRUNCATE(\'$(1)\')"
"SELECT * FROM customer WHERE {% trunc custname} = ?"

-- ifdef directive:
dbi.sqldirset.ifx.directive.ifdef.deflist = "INFORMIX"
dbi.sqldirset.ora.directive.ifdef.deflist = "ORACLE,ANSI"

"SELECT * FROM"
 || " {% ifdef INFORMIX "\"t1, OUTER(t2)\"}"
 || " {% ifdef ORACLE \"t1, t2\"}"
 || " WHERE t1.key = t2.col1 {% ifdef ORACLE \"(+)\"}"
```

To improve performances, a global switch could be used to disable SQL statement parsing for SQL
directives:

```
dbi.sql.directives = { true | false }
```

Starting with version 2.00, SQL directive sets are no longer supported. Consider writing SQL
statements with a syntax that is supported by all target database servers, or use
dynamic SQL if you need to adapt the SQL syntax to the database type.

## Related links

**Related concepts**  

[Dynamic SQL management](../10_sql-support/1140-dynamic-sql-management.md "Explains how to execute and manage SQL statements at runtime.")

[Static SQL cache is removed](0319-static-sql-cache-is-removed.md "The Static SQL Cache has been removed.")
