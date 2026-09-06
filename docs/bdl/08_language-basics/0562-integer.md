---
title: "INTEGER"
source: "fgl-topics/c_fgl_datatypes_INTEGER.html"
breadcrumb: "Language basics > Primitive Data types > INTEGER"
type: "concept"
---

# INTEGER

> The INTEGER data type is used for storing large whole numbers.

## Syntax

```
INTEGER
```

1. `INT` and `INTEGER` are synonyms.

## Usage

The storage of `INTEGER` variables is based on 4 bytes of signed data
( = 32 bits ).

The value range is from -2,147,483,647 to +2,147,483,647.

`INTEGER` variables can be initialized with [integer
literals](0586-integer-literals.md "Integer literals define a whole number in an expression."):

```
MAIN
  DEFINE i INTEGER
  LET i = 1234567
  DISPLAY i 
END MAIN
```

When assigning a whole number that exceeds the `INTEGER` range, the
overflow error [-1215](../15_library-reference/4483-genero-bdl-errors.md) will be raised.

`INTEGER` variables are initialized to zero in functions, modules and
globals.

The `INTEGER` type can be used to define variables storing values from
SERIAL columns.

Data type conversion can be controlled by catching the runtime exceptions. For more
details, see [Handling type conversion errors](0579-handling-type-conversion-errors.md "Runtime errors can be handled on type conversion failures.").

## Related links

**Related concepts**  

[BIGINT](0554-bigint.md "The BIGINT data type is used for storing very large whole numbers.")

[SMALLINT](0566-smallint.md "The SMALLINT data type is used for storing small whole numbers.")

[TINYINT](0568-tinyint.md "The TINYINT data type is used for storing very small whole numbers.")
