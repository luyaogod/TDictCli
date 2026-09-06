---
title: "fglhint_* in SQL comments"
source: "fgl-topics/c_fgl_Migrate_to_310_fglhints.html"
breadcrumb: "Upgrading > Upgrade Guides for Genero BDL > BDL 3.10 upgrade guide > fglhint_* in SQL comments"
type: "concept"
---

# fglhint_* in SQL comments

> Using SQL comment hints to control statement execution.

Starting with version 3.10, you can now specify `fglhint_*` keywords in C-style
SQL comments, to give an indication to the database driver about the type of SQL statement to be
executed.

> **Note:**
>
> C-style SQL comments can only be used in dynamic SQL statements and SQL blocks.

For example, to force an `INSERT` statement to be treated as a regular
`SELECT` returning a result set, use the `fglhint_select` hint in
a C-style comment:

```
DECLARE c1 CURSOR
   FROM "/* fglhint_select */ INSERT INTO table1 OUTPUT INSERTED.* SELECT * FROM customers"
```

> **Important:**
>
> If you are using `/* */` comments, these will now be parsed
> and any unknown keyword will be ignored: Comments such as `/* INSERT */` or
> `/* SELECT */` must be replaced by `/* fglhint_insert */` and
> `/* fglhint_select */` respectively.

Furthermore, Informix® emulation can be
disabled with the `fglhint_no_ifxemul` hint.

## Related links

**Related concepts**  

[fglhint\_\* SQL comments](../10_sql-support/1147-fglhint-sql-comments.md "Using special SQL comment hints to control statement execution.")

[Static SQL statements](../10_sql-support/1115-static-sql-statements.md "Describes static SQL statements supported in the language.")

[Dynamic SQL management](../10_sql-support/1140-dynamic-sql-management.md "Explains how to execute and manage SQL statements at runtime.")
