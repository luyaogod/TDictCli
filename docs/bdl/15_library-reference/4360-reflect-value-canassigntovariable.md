---
title: "reflect.Value.canAssignToVariable"
source: "fgl-topics/c_fgl_ext_reflect_Value_canAssignToVariable.html"
breadcrumb: "Library reference > Extension packages > The reflect package > The reflect.Value class > reflect.Value methods > reflect.Value.canAssignToVariable"
type: "concept"
---

# reflect.Value.canAssignToVariable

> Checks if this reflect.Value can be assigned to a variable.

## Syntax

```
canAssignToVariable(
     var any-type )
  RETURNS BOOLEAN
```

1. var is a program [variable](../08_language-basics/0686-variables.md "Explains how to define program variables.").
2. any-type can be of a [various kind of types](../09_advanced-features/0847-various-type-specification.md "Some Genero APIs use variant types for parameters or returns.").

## Usage

The `canAssignToVariable()` method returns `TRUE`, if the current
`reflect.Value` object references a value that can be assigned to the variable passed
as parameter.

For example, it is not possible to assign a structured `RECORD` to a variable
defined with a primitive type such as `STRING`.

This method is typically used to check that a target variable can be assigned with the [`assignToVariable()`](4359-reflect-value-assigntovariable.md "Assigns this reflect.Value to a variable.")
method.

## Example

```
IMPORT reflect
MAIN
    DEFINE r1, r2 RECORD pkey INTEGER END RECORD
    DEFINE val reflect.Value
    DEFINE s STRING
    LET val = reflect.Value.valueOf(r1)
    IF NOT val.canAssignToVariable(s) THEN
        DISPLAY "Cannot assign r1 to s..."
    END IF
    IF val.canAssignToVariable(r2) THEN
        DISPLAY "Assigning r1 to r2..."
        CALL val.assignToVariable(r2)
    END IF
END MAIN
```

Shows:

```
Cannot assign r1 to s...
Assigning r1 to r2...
```

## Related links

**Related concepts**  

[Data type conversion reference](../08_language-basics/0578-data-type-conversion-reference.md "This topic lists type conversion rules for all data types.")
