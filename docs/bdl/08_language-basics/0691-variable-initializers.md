---
title: "Variable initializers"
source: "fgl-topics/c_fgl_variables_009.html"
breadcrumb: "Language basics > Variables > Variable initializers"
type: "concept"
---

# Variable initializers

> Variables can be initialized in their definition.

## Syntax

A variable initializer is one of the following language elements, to be
specified after an equal sign of a variable definition in a [`DEFINE`](0688-define.md "The DEFINE instruction declares a program variable with a given type.") or [`VAR`](0689-var.md "The VAR instruction declares a program variable within a code block.")
instruction.

An initializer
is:

```
{ scalar-initializer
| record-initializer
| array-initializer
| dictionary-initializer
}
```

where scalar-initializer
is:

```
{ integer-literal
| numeric-literal
| text-literal
| mdy-date-literal
| datetime-literal
| interval-literal
| boolean-literal
| NULL
}
```

and record-initializer
is:

```
( identifier : initializer [,...] )
```

and array-initializer
is:

```
[ initializer [,...] ]
```

and dictionary-initializer
is:

```
( key : initializer [,...] )
```

1. integer-literal is an [whole number literal](0586-integer-literals.md "Integer literals define a whole number in an expression.").
2. numeric-literal is a [decimal number literal](0587-numeric-literals.md "Numeric literals define values with a decimal part in an expression.").
3. text-literal is a [character string literal](0588-text-literals.md "Text literals define a character string in an expression."), including [localized strings](../09_advanced-features/0901-localized-strings.md "Localized strings provide a means of writing applications in which the text of strings can be customized on site.").
4. mdy-date-literal is an [`MDY(mm,dd,yyyy)`
   specification](0589-mdy-m-d-y-literals.md "MDY() literals define a DATE literal in an expression.").
5. datetime-literal is a [datetime literal](0590-datetime-literals.md "Datetime literals define date/time value in an expression.").
6. interval-literal is an [interval literal](0591-interval-literals.md "Interval literals define an interval value in an expression.").
7. key identifies a [dictionary](0742-dictionaries.md "A dictionary holds an unordered collection of elements accessed by a key.") entry. It must be a string literal, and cannot be a localized string.
8. When using `NULL` as initializer, it is equivalent to [`INITIALIZE variable-name TO
   NULL`](0698-initialize.md "The INITIALIZE instruction initializes program variables with NULL or default values.").

## Variable initializer basics

To initialize a variable in the `DEFINE` instruction, add an equal sign followed
by an initializer literal:

```
DEFINE v INTEGER = -800
```

A variable initializer can not use expressions such as a value returned by a function, it must
always be a static literal.

The compiler produces the error [-6631](../15_library-reference/4483-genero-bdl-errors.md), if the type of the initializer and the type of the variable are
incompatible:

```
MAIN
    DEFINE v INT = "abc"
| incompatible types, found: 'CHAR', required: 'INTEGER'.
| See error number -6631.
END MAIN
```

Local function variables of the first function in a module will be initialized before any
`WHENEVER` statement.

See the [DEFINE syntax topic](0688-define.md "The DEFINE instruction declares a program variable with a given type.") for a complete
description of the variable initializer syntax.

Any variable of any type can be initialized with a simple `= NULL` initializer,
which is equivalent to the `INITIALIZE variable-name TO NULL`
instruction:

```
DEFINE rec RECORD f1 INT, f2 VARCHAR(20) END RECORD = NULL
DEFINE arr ARRAY[100] OF DATE = NULL
```

## Initializing DATE variables

Date variables can be initialized with a `MDY()`
literal:

```
DEFINE d DATE = MDY(12,24,2018)
```

## Initializing RECORD variables

Record variables can be initialized with a list of identifiers followed by a colon and the
literal value, surrounded with
parentheses:

```
DEFINE r1 RECORD
        f1 INTEGER,
        f2 STRING 
    END RECORD = (f1: 99, f2: "abc")
```

An invalid record initialization value can produce compilation errors such as [-8421](../15_library-reference/4483-genero-bdl-errors.md), [-8424](../15_library-reference/4483-genero-bdl-errors.md), or [-8425](../15_library-reference/4483-genero-bdl-errors.md).

Variables defined with [types](0696-user-defined-types.md "User defined types help to centralize the definition of complex structured data types.") can also be
initialized:

```
TYPE Type1 RECORD
        f1 INTEGER,
        f2 STRING 
    END RECORD
DEFINE r2 Type1 = (f1: 99, f2: "abc")
```

Record members can be omitted, or specified in a different
order:

```
DEFINE r1 Type1 = (f2: "abc")            -- f1 omitted
DEFINE r2 Type1 = (f2: "abc", f1: 99)    -- f2 before f1
```

When a record member is omitted in the initializer, it will get the default value according to
its type, as described in [Variable default values](0697-variable-default-values.md "Variables get a default value when defined.").

Variables defined from [database schema columns](0695-database-column-types.md "Simple variables and record structures can be defined from database columns types.") can
be initialized by specifying the initializer after the `LIKE`
clause:

```
SCHEMA stores
DEFINE customer RECORD LIKE customer.* = (fname: 'Peter', lname: 'Mango')
```

Initiliazers for nested records must specify the sub-record member name followed by parentheses
including the values for the members of the the
sub-record:

```
DEFINE r1 RECORD
          f1 INTEGER,
          f2 RECORD
             f21 STRING,
             f22 STRING
          END RECORD
    END RECORD = ( f1: 99,
                   f2: (f21:"abc", f22:"def") )
```

## Initializing DYNAMIC ARRAY variables

Arrays can be initialized by specifying a list of elements separated by a comma, surrounded by
square brackets:

```
DEFINE arr DYNAMIC ARRAY OF INT = [ 1, 2, 3, 5, 8, 13 ]
```

An invalid array initialization value can produce compilation errors such as [-8422](../15_library-reference/4483-genero-bdl-errors.md).

Structured arrays can be initialized by combining the initializer syntax for arrays and
records:

```
DEFINE pls DYNAMIC ARRAY OF RECORD
                id INTEGER,
                name VARCHAR(50)
           END RECORD =
         [
           ( id: 501, name: "Baxter" ),
           ( id: 502, name: "Folkap" ),
           ( id: 503, name: "Kirtshof" )
         ]
```

Multi-dimensional static array:

```
DEFINE matrix1 ARRAY[5,3] OF INTEGER = [
 [11,12,63],
 [58,23,63],
 [21,67,36],
 [11,65,31],
 [14,12,32]
]
```

## Initializing DICTIONARY variables

`DICTIONARY` variables can be initialized by using a comma-separated list of
elements included in parentheses. Each element must be a key/value pair separated by a
colon:

```
DEFINE mcs DICTIONARY OF DECIMAL(10,5) =
         ( "Pi"      : 3.14159,
           "Euler"   : 2.71828,
           "Golden"  : 1.61803
         )
```

Structured dictionaries can be initialized by combining the initializer syntax for dictionaries
and records:

```
DEFINE pls DICTIONARY OF RECORD
                name VARCHAR(50),
                address VARCHAR(200)
           END RECORD =
         (
           "CB841" : ( name: "Baxter", address : "4 Baker Street" ),
           "CB112" : ( name: "Folkap", address : "234 Sunset Road" ),
           "CB233" : ( name: "Kirtshof", address : "76 Colmor Row" )
         )
```

## Related links

**Related concepts**  

[Literals](0585-literals.md "Describes the syntax of literals (constant values) to be used in sources.")
