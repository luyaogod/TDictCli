---
title: "CONSTANT"
source: "fgl-topics/c_fgl_Constants_003.html"
breadcrumb: "Language basics > Constants > CONSTANT"
type: "concept"
---

# CONSTANT

> The CONSTANT instruction defines a program constant.

## Syntax

```
[PRIVATE|PUBLIC] CONSTANT
 identifier [ data-type ] = literal
 [,...]
```

1. identifier is the name of the constant to be defined, that must follow the
   convention for [identifiers](0551-identifiers.md "A Genero BDL identifier is a sequence of characters used to identify a program entity.").
2. data-type can be any [primitive data
   type](0553-primitive-data-types.md "Selecting the correct data type assists you in the input, storage, and display of your data."), except complex types like `TEXT` or `BYTE`.
3. literal must be an [integer](0586-integer-literals.md "Integer literals define a whole number in an expression."),
   [decimal](0587-numeric-literals.md "Numeric literals define values with a decimal part in an expression."), [string](0588-text-literals.md "Text literals define a character string in an expression."), or [date/time](0590-datetime-literals.md "Datetime literals define date/time value in an expression."), [interval](0591-interval-literals.md "Interval literals define an interval value in an expression.") literal, or an [`MDY()`](0589-mdy-m-d-y-literals.md "MDY() literals define a DATE literal in an expression.") expression.
4. literal cannot be `NULL`.

## Usage

Define constants to name static values that do not change during the program execution.

Constants are similar to variables, except that its value cannot be modified.

Constants can be defined with global, module, or function scope.

By default, module constants are private; they cannot be used by an other module of the
program. To make a module constant public, add the `PUBLIC` keyword before
`CONSTANT`. When a module constant is declared as public, it can be referenced by
another module by using the [`IMPORT`](../09_advanced-features/0813-importing-modules.md "Use the IMPORT ... instruction to import BDL, C or Java external modules in the current module.")
instruction.

When declaring a constant, the [data type](0553-primitive-data-types.md "Selecting the correct data type assists you in the input, storage, and display of your data.")
specification can be omitted. The [literal value](0585-literals.md "Describes the syntax of literals (constant values) to be used in sources.")
automatically defines the data
type:

```
CONSTANT c1 = "Drink" -- Declares a STRING constant 
CONSTANT c2 = 4711    -- Declares an INTEGER constant
```

However, in some cases, you may need to specify the exact data type for a
constant:

```
CONSTANT c1 SMALLINT = 12000 -- Would be an INTEGER by default
```

Constants can be used in variable, records, and array definitions:

```
CONSTANT n = 10
DEFINE a ARRAY[n] OF INTEGER
```

Constants can be used at any place in the language where you normally use
literals:

```
CONSTANT n = 10
FOR i=1 TO n
  ...
```

Constants can be passed as function parameters, and returned from functions.

Define public constants in a module to be imported by
others:

```
PUBLIC CONSTANT pi = 3.14159265
```

For date time constants, the value must be specified as an [`MDY()`](0589-mdy-m-d-y-literals.md "MDY() literals define a DATE literal in an expression.") literal, [`DATETIME` literal](0590-datetime-literals.md "Datetime literals define date/time value in an expression.") or [`INTERVAL`
literal](0591-interval-literals.md "Interval literals define an interval value in an expression."):

```
CONSTANT my_date DATE = MDY(12,24,2011)
CONSTANT my_datetime DATETIME YEAR TO SECOND
              = DATETIME(2011-12-24 11:22:33) YEAR TO SECOND
CONSTANT my_interval INTERVAL HOUR(5) TO FRACTION(3)
              = INTERVAL(-54351:50:24.234) HOUR(5) TO FRACTION(3)
```

A constant cannot be used in the `ORDER BY` clause of a static
`SELECT` statement, because the compiler considers identifiers after
`ORDER BY` as part of the SQL statement (i.e. column names), not as
constants:

```
CONSTANT pos = 3
-- Next line will produce an error at runtime
SELECT * FROM customers ORDER BY pos
```

Automatic data type conversion can take place in some
cases:

```
CONSTANT c1 CHAR(10) = "123"
CONSTANT c2 CHAR(10) = "abc"
DEFINE i INTEGER
FOR i = 1 TO c1 -- Constant "123" is converted to 123 integer
   ...
FOR i = 1 TO c2 -- Constant "abc" is converted to zero!
   ...
```

Character constants defined with a string literal that is longer than the length
of the data type are truncated:

```
CONSTANT s CHAR(3) = "abcdef"
DISPLAY s  -- Displays "abc"
```

The compiler throws an error when an undefined symbol is used in a constant
declaration:

```
CONSTANT s CHAR(c) = "abc"
-- Compiler error: c is not defined.
```

The compiler throws an error when a variable is used in a constant
declaration:

```
DEFINE c INTEGER
CONSTANT s CHAR(c) = "abc"
-- Compiler error: c is a variable, not a constant.
```

The compiler throws an error when you try to assign a value to a
constant:

```
CONSTANT c INTEGER = 123
LET c = 345
-- Runtime error: c is a constant.
```

The compiler throws an error when the symbol used is not defined as an integer
constant:

```
CONSTANT c CHAR(10) = "123"
DEFINE s CHAR(c)
-- Compiler error: c is a not an integer constant.
```

You typically define common special characters with
constants:

```
CONSTANT c_esc  = '\x1b'
CONSTANT c_tab  = '\t'
CONSTANT c_cr   = '\r'
CONSTANT c_lf   = '\n'
CONSTANT c_crlf = '\r\n'
```

## Related links

**Related concepts**  

[Importing modules](../09_advanced-features/0813-importing-modules.md "Use the IMPORT ... instruction to import BDL, C or Java external modules in the current module.")
