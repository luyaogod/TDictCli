---
title: "Returning values"
source: "fgl-topics/c_fgl_Functions_returns.html"
breadcrumb: "Language basics > Functions > Returning values"
type: "concept"
---

# Returning values

> A function can return values with the RETURN instruction.

## Defining the returned types in a function declaration

Function definitions can specify the list of data types returned by the function with the
`RETURNS` clause.

The `RETURNS` clause in the function header is not mandatory. However, it is
recommended, to define the complete function signature, and take advantage of better code checking
by the compiler.

If the function returns a single value, specify the data type after the `RETURNS`
clause:

```
FUNCTION get_count() RETURNS INTEGER
  ...
END FUNCTION
```

When the function returns a single value/type, it is also possible to enclose the type in
parentheses:

```
FUNCTION get_count() RETURNS (INTEGER)
```

But for convenience, the
parentheses are optional when only one return type is used.

If the function returns several values, you must specify the data types after the
`RETURNS` clause inside
parentheses:

```
FUNCTION get_address() RETURNS ( CHAR(5), VARCHAR(100), INTEGER )
  ...
END FUNCTION
```

When the function returns no values, enforce compiler verification by using the
`RETURNS()` clause in the function
head:

```
DEFINE the_list DYNAMIC ARRAY OF STRING

FUNCTION clear_list() RETURNS ()
    CALL the_list.clear()
    RETURN 0
| Invalid number of return values.
| See error number -8415.
END FUNCTION
```

The `RETURNS` types specification can use base data types such as
`INTEGER`, user-defined types, and
classes:

```
SCHEMA stores
TYPE t_cust DYNAMIC ARRAY OF RECORD LIKE customer.*
...
FUNCTION get_cust_list() RETURNS t_cust
  ...
END FUNCTION
```

## Returning from the function

Use the [`RETURN`](0676-return.md "The RETURN instruction gives the control of execution back to the caller, optionally returning values on the stack.") instruction in
the body of the function, to push a list of values on the stack, and return to the caller.

The `RETURN` instruction takes an optional comma-separated list of
expressions:

```
FUNCTION get_address() RETURNS ( CHAR(5), VARCHAR(100), INTEGER )
  DEFINE zipcode CHAR(5),
         street VARCHAR(100),
         city INTEGER
  ...
  RETURN zipcode, street, city
END FUNCTION
```

The next example shows a function returning a single
value:

```
FUNCTION get_count() RETURNS INTEGER
   ...
   RETURN count
END FUNCTION
```

If a function does not need to return a value, the `RETURN` instruction can be
used without arguments:

```
FUNCTION show_notfound() RETURNS ()
   IF sqlca.sqlcode==0 THEN
      RETURN
   END IF
   ...
END FUNCTION
```

## Expressions using functions returning a single value

When a function returns a single value, it can be invoked in [expressions](0592-expressions.md "Shows the possible expressions supported in the language."):

```
MAIN
    DEFINE r INTEGER
    LET r = 500 + add(50,5)
END MAIN

FUNCTION add(x,y) RETURNS INTEGER
    DEFINE x, y INTEGER
    RETURN (x + y)
END FUNCTION
```

## Calling a function returning a list of values

Functions returning multiple values must be invoked with a `CALL` instruction
using the `RETURNING` clause. Values specified in `RETURN` statement
must correspond in number and position to the `RETURNING` clause of the
`CALL` instruction, and must be of the same or of compatible data types, to the
variables in the `RETURNING` clause of the `CALL` statement. An error
results if the list of returned values in the `RETURN` statement conflicts in number
or in data type with the `RETURNING` clause of the `CALL` statement
that invokes the
function.

```
MAIN
  DEFINE zipcode CHAR(5),
         street VARCHAR(100),
         city VARCHAR(50)
  CALL get_address() RETURNING zipcode, street, city
END MAIN

FUNCTION get_default_address() RETURNS (STRING,STRING,STRING)
   RETURN "00000", "<undefined>", "<undefined>"
END FUNCTION
```

## Type casting of returned values

When using the `RETURNS` clause in a function definition, the data type(s) of the
returned value(s) will match the function definition.

In the next example, the function definition specifies `RETURNS INTEGER`.
Consequently, the arithmetic division will return a whole number. In another context, that numeric
expression would produce a decimal value:

```
MAIN
    DISPLAY divide( 5, 3 )
    DISPLAY ( 5 / 3 )
END MAIN

FUNCTION divide(
     dividend INTEGER,
     divisor INTEGER
) RETURNS INTEGER
    RETURN ( dividend / divisor )
END FUNCTION
```

Output:

```
          1
1.666666666666666666667
```

This matters for example when you create a variable from the value returned by a
function:

```
VAR v1 = divide(5,3)    -- v1 becomes an INTEGER
```

## Returning complex structures

When returning simple built-in types like `INTEGER`, values are copied on the
stack and copied to the caller variables.

To return a a [`RECORD`](0715-records.md "Records allow structured program variables definitions.") structure, use a
[fully
typed function definition](0763-function-definitions.md), and specify the record name (without `.*` notation)
in the `RETURNS`
instruction:

```
TYPE t_cust RECORD
         cust_id INTEGER,
         cust_name VARCHAR(50)
     END RECORD

MAIN
    DEFINE r1 t_cust
    LET r1 = get_customer()
    DISPLAY r1.*
END MAIN

FUNCTION get_customer() RETURNS (t_cust)
    DEFINE r t_cust
    LET r.cust_id = 999
    LET r.cust_name = "Mike"
    RETURN r
END FUNCTION
```

Records returned from functions as single structured value can be used in expressions and passed
directly to other functions, as in the following
example:

```
TYPE t_cust RECORD
         cust_id INTEGER,
         cust_name VARCHAR(50)
     END RECORD

MAIN
    DEFINE r1 t_cust
    LET r1 = init_customer( new_customer() )
    DISPLAY r1.*
END MAIN

FUNCTION new_customer() RETURNS (t_cust)
    DEFINE r t_cust
    LET r.cust_id = -1
    LET r.cust_name = NULL
    RETURN r
END FUNCTION

FUNCTION init_customer(r t_cust) RETURNS (t_cust)
    IF r.cust_name IS NULL THEN
        LET r.cust_name = "<undefined>"
    END IF
    RETURN r
END FUNCTION
```

When returning complex variables such class objects or dynamic arrays, the reference of the
element are copied on the stack (this means that you can create an object inside a function, and
return its reference in the `RETURN`
statement):

```
MAIN
    DEFINE c base.Channel
    LET c = open_file("myfile.txt")
    ...
END MAIN

FUNCTION open_file(filename)
    DEFINE filename STRING
    DEFINE c base.Channel
    TRY
        LET c = base.Channel.create()
        CALL c.openFile(filename, "r")
        RETURN c
    CATCH
        RETURN NULL
    END TRY
END FUNCTION
```

## Related links

**Related concepts**  

[Type conversions](0576-type-conversions.md "Explains primitive data type conversion rules of the language.")

[Runtime stack](../09_advanced-features/0833-runtime-stack.md "The runtime stack is used to pass/return values to/from functions.")
