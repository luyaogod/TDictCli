---
title: "SQL cursor management"
source: "fgl-topics/c_fgl_MigI4GL_031.html"
breadcrumb: "Upgrading > Migrating from IBM® Informix® 4GL to Genero BDL > 4GL programming topics > SQL cursor management"
type: "concept"
---

# SQL cursor management

> This topic describes differences between I4GL and Genero BDL in SQL cursor management.

## OPEN USING followed by FOREACH

In earlier versions of IBM® Informix® 4GL (I4GL), the `FOREACH` instruction had no
`USING` clause to pass SQL parameters to the prepared statement. SQL Parameters could
be specified in a `OPEN cursor-name USING
variable-list` instruction preceding the `FOREACH`
instruction:

```
PREPARE st1 FROM "SELECT * FROM tab WHERE col>?"
DECLARE cu1 CURSOR FOR st1
OPEN cu1 USING var 
FOREACH cu1 INTO rec.*
  DISPLAY rec.*
END FOREACH
```

This feature is supported by Genero Business Development Language, but can lead to defects with
some versions of the Informix database client. Review
your code to avoid the `OPEN` statement by moving the `USING` clause
to the `FOREACH` instruction.

> **Important:**
>
> Older versions of Informix 4GL did not support the `USING` clause in the
> `FOREACH` instruction. The solution was to issue an [`OPEN`](../10_sql-support/1151-open-result-set-cursor.md "Executes the SQL statement with result set associated with the specified database cursor") instruction with a
> `USING` clause, before the `FOREACH` (without `USING`
> clause). This practice is allowed by the Genero compiler for backward compatibility, but can lead to
> unexpected behavior, especially when enabling Informix optimization options such as OPTOFC=1. When
> the SQL statement has parameters, specify the parameter variables in a `USING` clause
> of the `FOREACH` instruction, and remove the unnecessary `OPEN`
> instruction before the `FOREACH`.

## **DECLARE might throw unexpected errors**

With IBM Informix 4GL, a `DECLARE ... CURSOR` instrucion does not throw an SQL error
when using a table that does not exist. With Genero BDL, the `DECLARE` instruction
will fail:

```
DECLARE c1 CURSOR FOR SELECT * FROM tt1
CREATE TEMP TABLE tt1 ( pkey INTEGER, name VARCHAR(50) )
```

Such code needs to be fixed: An SQL table must exist before it is used.

That programming mistake can typically occur, when the cursor declaration and table creation are
encapsulated in functions:

```
MAIN
    DATABASE test1
    CALL declare_cursor()
    CALL create_table()
    CALL open_cursor()
END MAIN

FUNCTION create_table()
    CREATE TEMP TABLE tt1 ( pkey INTEGER, name VARCHAR(50) )
END FUNCTION

FUNCTION declare_cursor()
    DECLARE c1 CURSOR FOR SELECT * FROM tt1
END FUNCTION

FUNCTION open_cursor()
    OPEN c1
END FUNCTION
```

To fix the above code, switch the lines calling the `create_table()` and
`declare_cursor()` functions.

## FREE before DECLARE

The IBM Informix
4GL compiler accepts the `FREE` instrucion before the `DECLARE`
instruction, for example:

```
FREE c_cust
DECLARE c_ust CURSOR FOR SELECT * FROM customer
```

With Genero BDL, the above code will produce a compiler error.

This use case is not common and must be avoided. The `DECLARE` cursor instruction
must appear first in the module, before any other cursor instruction such as `OPEN`,
`CLOSE`, `FETCH`.

The typical case that reveals this difference is when cursor instructions are encapsulated in
functions, and the functions are declared in an unexpected order, as in the following
example:

```
FUNCTION c_cust_open()
    OPEN c_cust
END FUNCTION

FUNCTION c_cust_close()
    CLOSE c_cust
END FUNCTION

FUNCTION c_cust_free()
    FREE c_cust
END FUNCTION

FUNCTION c_cust_declare()
    DECLARE c_cust CURSOR FOR SELECT * FROM customer
END FUNCTION
```

Such source code must be fixed to move the `c_cust_declare` function at the
top.

## The CURSOR\_NAME() function

With IBM Informix 4GL, the `CURSOR_NAME()` function is available, to map the 4GL cursor
name to the the underlying cursor identifier procuded by the compiler, which is the real cursor id
used by the Informix database server. This function is typically used in `UPDATE / DELETE
WHERE CURRENT OF`:

```
MAIN
    DEFINE cn VARCHAR(50)
    DATABASE test1
    DECLARE mycursor CURSOR FOR SELECT * FROM customer
    LET cn = CURSOR_NAME("mycursor")
    DISPLAY cn
END MAIN
```

See Informix 4GL documentation for more details.

Genero BDL does not support the `CURSOR_NAME()` function: It is useless, because
the runtime mangles cursor names automatically.

The scope of an SQL cursor name or prepared statement name is the module. The underlying database
client API does not know about modules: The scope of cursor names is the SLQ connection. While I4GL
compiler mangles cursor names and statement names at compile time, FGL handles this by mapping
module cursor names to unique program-wide cursor names at runtime.

Note that if you do not use `IMPORT FGL` and strick `-Wimplicit`
compilation options of Genero fglcomp, since `CURSOR_NAME()` is a
function, it will not be detected at compile time, and it usage will only be detected with a runtime
error when this function is called.

Scan the sources to check for CURSOR\_NAME() usage, and if code change is not possible, write a
user function that directly returns the cursor name passed as
parameter:

```
FUNCTION cursor_name(name STRING) RETURNS STRING
    RETURN name
END FUNCTION
```

## Related links

**Related concepts**  

[Result set processing](../10_sql-support/1148-result-set-processing.md "Shows how to fetch rows from a database query.")
