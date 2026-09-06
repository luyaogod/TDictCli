---
title: "Stored functions defined with output parameters"
source: "fgl-topics/c_fgl_sql_programming_012.html"
breadcrumb: "SQL support > SQL database guides > IBM® Informix® > Fully supported IBM® Informix® SQL features > Stored procedure calls > Stored functions defined with output parameters"
type: "concept"
description: "Starting with IDS 10.00, IBM® Informix® introduced the concept of output parameters for stored functions. To retrieve the output parameters, you must execute the routine in a SELECT statement defining ..."
---

# Stored functions defined with output parameters

Starting with IDS 10.00, IBM®
Informix® introduced the concept of output parameters for
stored functions.

To retrieve the output parameters, you must execute the routine in a `SELECT`
statement defining Statement Local Variables. These variables will be listed in the
select clause to be fetched as regular column values. See Informix documentation for more details.

In order to retrieve returning values into program variables, use an `INTO` clause
in the `EXECUTE` instruction.

This example shows how to call a stored function with output
parameters:

```
MAIN
    DEFINE pi, pr INTEGER
    DATABASE test1
    EXECUTE IMMEDIATE "create function proc2(i INT, OUT r INT)"
                    ||" returning int;"
                    ||" let r=i+10;"
                    ||" return 1;"
                    ||" end function"
    PREPARE s FROM "select r from systables where tabid=1 and proc2(?,r#int)==1"
    LET pi = 33
    EXECUTE s USING pi INTO pr 
    DISPLAY "Output value: ", pr 
    EXECUTE IMMEDIATE "drop function proc2"
END MAIN
```
