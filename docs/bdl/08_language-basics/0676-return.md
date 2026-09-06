---
title: "RETURN"
source: "fgl-topics/c_fgl_FlowControl_RETURN.html"
breadcrumb: "Language basics > Flow control > RETURN"
type: "concept"
---

# RETURN

> The RETURN instruction gives the control of execution back to the caller, optionally returning values on the stack.

## Syntax

```
RETURN [ value [,...]  ]
```

1. value can be any valid expression, an object reference or complex type
   reference such as a dynamic array reference.

## Usage

The `RETURN` instruction transfers the control back from a [function](0761-functions.md "Describes user defined functions.") with optional return values.

[Record](0715-records.md "Records allow structured program variables definitions.") members can be returned with the
`.*` or `THRU` notation. Each member is returned as an independent
variable.

Consider using the [fully typed function](0763-function-definitions.md "A FUNCTION definition defines a named procedure with a set of statements.")
definition syntax, with a `RETURNS` clause in the function header, to get better
compilation control of your code. When using the `RETURNS` clause, the compiler will
check that the function body contains `RETURN` instructions that match the number of
return values as specified in the function definition.

A function may have several `RETURN` points (not recommended in structured
programming) but they must all return the same number of values.

The number of returned values must correspond to the number of [variables](0686-variables.md "Explains how to define program variables.") listed in the `RETURNING` clause of
the `CALL` statement invoking this function.

A function cannot return a [static array](0732-static-arrays.md "Static arrays have a predefined and limited size."), but can
return the reference of a [dynamic array](0734-dynamic-arrays.md) or [dictionary](0742-dictionaries.md "A dictionary holds an unordered collection of elements accessed by a key.").

## Example

```
MAIN
  DEFINE fname, lname VARCHAR(30)
  LET lname = lastname(943)
  DISPLAY lname
  CALL fullname(235) RETURNING fname, lname
  DISPLAY fname, lname
END MAIN

FUNCTION lastname(id INTEGER) RETURNS STRING
  CASE id
    WHEN 943
      RETURN "McTiger"
    OTHERWISE
      RETURN NULL
  END CASE
END FUNCTION

FUNCTION fullname(id INTEGER) RETURNS(STRING, STRING)
  CASE id
    WHEN 235
      RETURN "Lee", "Park"
    OTHERWISE
      RETURN NULL, NULL
  END CASE
END FUNCTION
```

## Related links

**Related concepts**  

[Returning values](0768-returning-values.md "A function can return values with the RETURN instruction.")

[CALL](0675-call.md "The CALL instruction invokes a specified function or method.")

[Expressions](0592-expressions.md "Shows the possible expressions supported in the language.")

[Type conversions](0576-type-conversions.md "Explains primitive data type conversion rules of the language.")
