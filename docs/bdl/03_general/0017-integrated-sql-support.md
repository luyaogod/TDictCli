---
title: "Integrated SQL support"
source: "fgl-topics/c_fgl_intro_BDL_010.html"
breadcrumb: "General > Introduction to Genero BDL programming > Genero BDL concepts > Integrated SQL support"
type: "concept"
---

# Integrated SQL support

> A set of SQL statements are part of the language syntax and can be used directly in the source code, as a normal procedural instruction.

The static SQL statements are parsed and validated at compile time. At runtime, these SQL
statements are automatically prepared and executed. Program variables are detected by the
compiler and handled as SQL parameters. Common SQL statements such as SELECT, INSERT, UPDATE
or DELETE can be directly written in the source code, as part of the language syntax:

```
MAIN
  DEFINE n INTEGER, s CHAR(20)
  DATABASE stores
  LET s = "Sansino"
  SELECT COUNT(*) INTO n FROM customer WHERE custname = s 
  DISPLAY "Rows found: " || n 
END MAIN
```

Dynamic SQL management allows you to execute SQL statements that are constructed at runtime. The
SQL statement can use SQL parameters:

```
MAIN
  DEFINE txt CHAR(20)
  DATABASE stores 
  LET txt = "SET DATE_FORMAT = YMD"
  PREPARE sh FROM txt 
  EXECUTE sh 
END MAIN
```

Through the database drivers, the same program can open database
connections to any of the supported databases.

## Related links

**Related concepts**  

[SQL support](../10_sql-support/0983-sql-support.md "These topics cover SQL support in the Genero Business Development Language.")
