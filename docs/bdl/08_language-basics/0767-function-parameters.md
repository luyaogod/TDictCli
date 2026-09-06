---
title: "Function parameters"
source: "fgl-topics/c_fgl_Functions_parameters.html"
breadcrumb: "Language basics > Functions > Function parameters"
type: "concept"
---

# Function parameters

> Functions can take parameters, to specialize their behavior.

## Purpose of function parameters

The function declaration specifies the name of the function and the identifiers of its formal
arguments (if any).

Function parameters hold values passed by the caller, to be used in the body of the function.

## Function parameter specification with legacy syntax

With the legacy function definition syntax, the data type of each formal argument of the function
must be specified by a `DEFINE` statement that immediately follows the argument
list.

```
FUNCTION check_address(zipcode, street, city)
  DEFINE zipcode CHAR(5),
         street VARCHAR(100),
         city VARCHAR(50)
  DEFINE found BOOLEAN -- local function variable
  ...
END FUNCTION
```

## Function parameter specification with fully-typed syntax

With the fully-typed function definition syntax, the type of the parameters is specified inside
the parentheses:

```
FUNCTION check_address(zipcode CHAR(5), street VARCHAR(100), city VARCHAR(50))
  DEFINE found BOOLEAN -- local function variable
  ...
END FUNCTION
```

## Function without parameters

If no argument is needed in a function call, an empty argument list must still be supplied,
enclosed between the
parentheses:

```
FUNCTION begin_work()
   ...
END FUNCTION
```

## Passing primitive type values as parameters

Function arguments using simple data types such as [`INTEGER`](0562-integer.md "The INTEGER data type is used for storing large whole numbers.") are passed by value (the values are copied on the stack).

In the following code example, the variable `x` defined in the
`MAIN` block will not be modified by the function:

```
MAIN
  DEFINE x INTEGER
  LET x = 123
  CALL myfunc(x)
  DISPLAY x   -- displays 123
END MAIN

FUNCTION myfunc(x)
  DEFINE x INTEGER
  LET x = x + 1
END FUNCTION
```

Parameters of type [`TEXT`](0569-text.md "The TEXT data type stores large text data.") and [`BYTE`](0555-byte.md "The BYTE data type stores any type of binary data, such as images or sounds.") are also passed by value, but it's
only the large object handler which is copied on the stack. The actual data is pointed to by the
handler and therefore it is not copied.

## Type conversions

The actual argument in a call to the function need not be of the declared data type of the formal
argument. The runtime system will do the appropriate data type conversions when needed. If the data
type conversion is not possible, a runtime error occurs.

Genero BDL data type conversion is quite permissive by default. The following example will not
raise an error, because the string value is silently converted to a `NULL
INTEGER`:

```
MAIN
    CALL add("aaaaa")
END MAIN

FUNCTION add(x)
    DEFINE x INTEGER
    DISPLAY "x = ", x
END FUNCTION
```

## Passing complex objects as parameters

Complex structures such as [dynamic arrays](0734-dynamic-arrays.md) and [objects](../09_advanced-features/0942-oop-support.md "Describes Object Oriented Programming basics in the language.") are passed by reference: The reference of the complex
object is copied on the stack and can be used as is inside the
function:

```
MAIN
    DEFINE ch base.Channel
    LET ch = base.Channel.create()
    CALL ch.openFile(arg_val(1), "r")
    CALL read_lines(ch)
END MAIN

FUNCTION read_lines(ch base.Channel)
    DEFINE s STRING
    WHILE TRUE
        LET s = ch.readLine()
        IF ch.isEOF() THEN EXIT WHILE END IF
        DISPLAY s
    END WHILE
END FUNCTION
```

## Passing records as parameters

When using the `.*` (dot star) notation, [records](0715-records.md "Records allow structured program variables definitions.") are expanded on the stack: Each field of the record is passed by value like a
primitive data type. This is supported for backward
compatibility:

```
CALL show_cust_info( r_cust.* )
```

To improve code robustness, consider using [fully typed
function definitions](0763-function-definitions.md), and pass records without the `.*` notation as a single
structured value. The record is still copied on the stask (the record of the caller cannot by
modified by the
function):

```
SCHEMA stores
TYPE t_cust RECORD LIKE customer.*
FUNCTION insert_cust(r t_cust) RETURNS ()
    INSERT INTO customers VALUES (r.*)
END FUNCTION
```

Records can also be passed by reference, when specifying the `INOUT` keyword in
the function
definition:

```
SCHEMA stores
TYPE t_cust RECORD LIKE customer.*
FUNCTION init_cust(r t_cust INOUT) RETURNS ()
    INITIALIZE r.* TO NULL
    LET r.customer_num = 0
    LET r.lname = "<undefined>"
END FUNCTION
```

> **Important:**
>
> Passing records by reference with `INOUT` is only supported
> when the called function is defined locally, or is known by the compiler because it was imported
> with `IMPORT FGL`.

For more details, see [Passing records as parameter](../09_advanced-features/0835-passing-records-as-parameter.md).

## Related links

**Related concepts**  

[Primitive Data types](0553-primitive-data-types.md "Selecting the correct data type assists you in the input, storage, and display of your data.")

[Variables](0686-variables.md "Explains how to define program variables.")
