---
title: "Passing records by reference"
source: "fgl-topics/c_fgl_optimization_016.html"
breadcrumb: "Advanced features > Optimization > Optimize your programs > Passing records by reference"
type: "concept"
description: "With CALL func( record-name .*) , RECORD structures are passed to functions by copying all record members on the stack. All record members are also copied on the stack, when returning a complete ..."
---

# Passing records by reference

With `CALL func(record-name.*)`, `RECORD`
structures are passed to functions by copying all record members on the stack.

All record members are also copied on the stack, when returning a complete record from a function
with `RETURN record-name.*`

If the record is big, this can lead to performances issues when the function is called many
times, or when chaining function calls passing complete records to each other. This is even more
resource consuming, when using [`CHAR()`
types](0975-passing-char-parameters-to-functions.md).

To improve performances, you can pass records by reference to functions: To pass the reference
of a record to a function, define the function with the `INOUT` keyword for the
record parameter. With this solution, records of the caller can be modified in the function (there
is no need to return the record). When calling the function, only specify the record name.

> **Tip:**
>
> Consider using user-defined types with the [`TYPE` instruction](../08_language-basics/0751-types.md "Types can be defined by the programmer to centralize the definition of complex/structured variables."), as well as [fully typed
> function](../08_language-basics/0763-function-definitions.md) definitions.

The next code example shows the execution time improvement when passing a record by
reference:

```
TYPE t_cust RECORD
         cust_id INTEGER,
         cust_name CHAR(100),
         cust_addr CHAR(2000)
     END RECORD

MAIN
    CONSTANT MAXCNT = 1000000
    DEFINE cust t_cust
    DEFINE ts DATETIME HOUR TO FRACTION(5)
    DEFINE x INTEGER

    LET cust.cust_id = 101
    LET cust.cust_name = "xxx"
    LET cust.cust_addr = "zzz"

    LET ts = CURRENT
    FOR x=1 TO MAXCNT
        CALL func1( cust.* ) RETURNING cust.*
    END FOR
    DISPLAY CURRENT HOUR TO FRACTION(5) - ts

    LET ts = CURRENT
    FOR x=1 TO MAXCNT
        CALL func2( cust )
    END FOR
    DISPLAY CURRENT HOUR TO FRACTION(5) - ts

END MAIN

FUNCTION func1( rec )
    DEFINE rec t_cust
    CALL func11( rec.* ) RETURNING rec.*
    RETURN rec.*
END FUNCTION

FUNCTION func11( rec )
    DEFINE rec t_cust
    RETURN rec.*
END FUNCTION

FUNCTION func2( rec t_cust INOUT ) RETURNS ()
    CALL func22(rec)
END FUNCTION

FUNCTION func22( rec t_cust INOUT ) RETURNS ()
END FUNCTION
```

Output:

```
  0:00:01.01965
  0:00:00.11381
```

For more details, see [Passing records as parameters](../08_language-basics/0767-function-parameters.md)
