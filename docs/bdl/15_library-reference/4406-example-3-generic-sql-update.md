---
title: "Example 3: Generic SQL UPDATE"
source: "fgl-topics/c_fgl_ext_reflect_example_3.html"
breadcrumb: "Library reference > Extension packages > The reflect package > Examples > Example 3: Generic SQL UPDATE"
type: "concept"
description: "This example shows how to set an SQL parameter for base.SqlHandle , with the reflection API method reflect.Value.getCurrentValue() . genupd.4gl source: IMPORT reflect MAIN DEFINE v_decimal ..."
---

# Example 3: Generic SQL UPDATE

This example shows how to set an SQL parameter for [`base.SqlHandle`](3015-the-sqlhandle-class.md "The base.SqlHandle class is a built-in class providing an API to execute parameterized SQL statements, with or without result sets."), with the reflection API method [`reflect.Value.getCurrentValue()`](4364-reflect-value-getcurrentvalue.md "Returns the actual (typed) value of this value object.").

`genupd.4gl` source:

```
IMPORT reflect
  
MAIN
    DEFINE v_decimal DECIMAL(10,2)
    DEFINE v_money MONEY(10,2)
    DEFINE v reflect.Value
    DEFINE sqlcmd STRING

    DATABASE test1

    CREATE TEMP TABLE tt1 ( pkey INT, c_decimal DECIMAL(10,2) )
    LET v_decimal = -9999.99
    INSERT INTO tt1 VALUES ( 101, v_decimal )

    LET sqlcmd = "UPDATE tt1 SET c_decimal = ? WHERE pkey = ?"
    LET v = reflect.Value.valueOf(v_decimal)
    CALL generic_update(sqlcmd, 101, v)

    LET sqlcmd = "UPDATE tt1 SET c_decimal = ? WHERE pkey = ?"
    LET v_money = v_decimal
    LET v = reflect.Value.valueOf(v_money)
    CALL generic_update(sqlcmd, 101, v)
    -- Using the toString() method will fail with some types because of
    -- unsupported conversions.
    CALL failing_generic_update(sqlcmd, 101, v)

END MAIN

FUNCTION generic_update(sqlcmd STRING, pkey INTEGER, var reflect.Value)
    DEFINE h base.SqlHandle
    LET h = base.SqlHandle.create()
    CALL h.prepare(sqlcmd)
    DISPLAY "Current value = ", var.getCurrentValue()
    CALL h.setParameter( 1, var.getCurrentValue() )
    CALL h.setParameter( 2, pkey )
    CALL h.execute()
END FUNCTION

FUNCTION failing_generic_update(sqlcmd STRING, pkey INTEGER, var reflect.Value)
    DEFINE h base.SqlHandle
    LET h = base.SqlHandle.create()
    CALL h.prepare(sqlcmd)
    CALL h.setParameter( 1, var.toString() )
    CALL h.setParameter( 2, pkey )
    CALL h.execute()
END FUNCTION
```

Output:

```
Current value =     -9999.99
Current value =     $-9999.99
Program stopped at 'genupd.4gl', line number 45.
SQL statement error number -1213.
A character to numeric conversion process failed
```
