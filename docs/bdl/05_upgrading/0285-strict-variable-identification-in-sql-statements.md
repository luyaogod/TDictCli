---
title: "Strict variable identification in SQL statements"
source: "fgl-topics/c_fgl_Migrate_to_220_strict_sql_variables.html"
breadcrumb: "Upgrading > Upgrade Guides for Genero BDL > BDL 2.20 upgrade guide > Strict variable identification in SQL statements"
type: "concept"
---

# Strict variable identification in SQL statements

> Program variable identification in static SQL statements is more strict in version 2.20 than older versions.

If you define a variable with the same name as an SQL object (i.e. table name, table alias), the
fglcomp compiler will raise an error because it will consider the program variable first. For
example, if the variable name matches the table or alias identifier, using
table.column in the SQL statement will be resolved as
variable.member, which does not exist.

The next code example will not compile because the program defines a variable using the same name
as the table alias `c`:

```
MAIN
   DEFINE c INTEGER
   DATABASE stores
   SELECT COUNT(*) INTO c
      FROM customer c
         WHERE c.fname IS NULL
| The variable 'c' is not a record.  It cannot reference record elements.
| See error number -4347.
END MAIN
```

The code also fails to compile with IBM®
Informix® 4GL 7.32, but it did compile with
version of Genero Business Development Language.

To work around this, you must either rename the program variable, or explicitly identify SQL
objects with the @ prefix in the SQL
statement:

```
MAIN
   DEFINE c INTEGER
   DATABASE stores
   SELECT COUNT(*) INTO c
      FROM customer c
         WHERE @c.fname IS NULL
END MAIN
```

Recompile all your programs to find the conflicts.
