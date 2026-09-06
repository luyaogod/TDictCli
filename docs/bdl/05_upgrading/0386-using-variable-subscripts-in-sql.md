---
title: "Using variable subscripts in SQL"
source: "fgl-topics/c_fgl_MigI4GL_036.html"
breadcrumb: "Upgrading > Migrating from IBM® Informix® 4GL to Genero BDL > 4GL programming topics > Using variable subscripts in SQL"
type: "concept"
---

# Using variable subscripts in SQL

> Use of subscript operations on host variables in static SQL statements is not supported by Genero BDL.

IBM® Informix® 4GL
allows to use subscripts operators on host variables in static SQL statements. For example:

```
DEFINE name VARCHAR(40)
SELECT * FROM tab WHERE col = name[2,3]
```

When doing this, I4GL will silently ignore the `[2,3]` operator, and will use the
full variable value!

Using this syntax is bad practice.

To avoid mistakes, the Genero BDL compiler denies this syntax by producing error
-4402:

```
MAIN
    DEFINE p_rec RECORD
               pkey INTEGER,
               name VARCHAR(10)
           END RECORD

    DATABASE test1

    CREATE TEMP TABLE tt1 ( pkey INTEGER, name VARCHAR(10) )
    INSERT INTO tt1 VALUES ( 101, 'ABC' )

    LET p_rec.pkey = 101
    LET p_rec.name = "XYZ"

    UPDATE tt1
       SET name = p_rec.name[1,1]
     WHERE pkey = p_rec.pkey

    DECLARE c0 CURSOR FOR SELECT * FROM tt1
    FOREACH c0 INTO p_rec.*
        DISPLAY p_rec.*
    END FOREACH

    SELECT pkey INTO p_rec.pkey FROM tt1 WHERE name = p_rec.name[1,1]
    DISPLAY status, p_rec.pkey

    PREPARE s1 FROM "SELECT pkey FROM tt1 WHERE name = ?"
    DECLARE c1 CURSOR FOR s1
    OPEN c1 USING p_rec.name[1,1]
    FETCH c1 INTO p_rec.pkey
    CLOSE c1

END MAIN
```

Compiler error messages:

```
$ fglcomp -M main.4gl
main.4gl:16:19:16:33:error:(-4402) In this type of statement, subscripting may ...
main.4gl:24:55:24:69:error:(-4402) In this type of statement, subscripting may ...
main.4gl:29:19:29:33:error:(-4402) In this type of statement, subscripting may ...
```

## Related links

**Related concepts**  

[Using program variables in static SQL](../10_sql-support/1117-using-program-variables-in-static-sql.md "Static SQL syntax supports the usage of program variables as SQL parameters.")
