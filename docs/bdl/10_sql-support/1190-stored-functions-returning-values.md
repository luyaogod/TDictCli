---
title: "Stored functions returning values"
source: "fgl-topics/c_fgl_sql_programming_011.html"
breadcrumb: "SQL support > SQL database guides > IBM® Informix® > Fully supported IBM® Informix® SQL features > Stored procedure calls > Stored functions returning values"
type: "concept"
description: "To return values from an IBM® Informix® SPL routine, execute the routine and fetch the output values, as you would for a regular SELECT statement producing a result set. Note: Informix distinguishes ..."
---

# Stored functions returning values

To return values from an IBM®
Informix® SPL routine, execute the routine and fetch the
output values, as you would for a regular `SELECT` statement producing a result
set.

> **Note:**
>
> Informix distinguishes between
> stored functions from stored procedures. Only stored functions (with a
> `RETURNING` clause) can return values. Stored procedures do not return values.

To execute an Informix stored function
from a BDL program, use the `EXECUTE FUNCTION` SQL
instruction:

```
PREPARE stmt FROM "execute function proc1(?)"
```

In order to retrieve returning values into program variables, use an `INTO` clause
in the `EXECUTE` instruction.

This example shows how to call a stored
function:

```
MAIN
   DEFINE n INTEGER
   DEFINE d DECIMAL(6,2)
   DEFINE c VARCHAR(200)
   DATABASE test1
   EXECUTE IMMEDIATE "create function proc1( p1 integer )"
                || " returning decimal(6,2), varchar(200);"
                || "  define p2 decimal(6,2);"
                || "  define p3 varchar(200);"
                || "  let p2 = p1 + 0.23;"
                || "  let p3 = 'Value = ' || p1;"
                || "  return p2, p3;"
                || " end function;"
   PREPARE stmt FROM "execute function proc1(?)"
   LET n = 111
   EXECUTE stmt USING n INTO d, c
   DISPLAY d
   DISPLAY c
END MAIN
```
