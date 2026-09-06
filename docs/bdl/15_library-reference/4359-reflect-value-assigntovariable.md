---
title: "reflect.Value.assignToVariable"
source: "fgl-topics/c_fgl_ext_reflect_Value_assignToVariable.html"
breadcrumb: "Library reference > Extension packages > The reflect package > The reflect.Value class > reflect.Value methods > reflect.Value.assignToVariable"
type: "concept"
---

# reflect.Value.assignToVariable

> Assigns this reflect.Value to a variable.

## Syntax

```
assignToVariable( var any-type )
```

1. var is a program [variable](../08_language-basics/0686-variables.md "Explains how to define program variables.").
2. any-type can be of a [various kind of types](../09_advanced-features/0847-various-type-specification.md "Some Genero APIs use variant types for parameters or returns.").

## Usage

The `assignToVariable()` method sets the variable passed as parameter with the
value of the current `reflect.Value` object.

Before calling this method, you can check if a variable can be assigned with this
`reflect.Value` object by using the [`canAssignToVariable()`](4360-reflect-value-canassigntovariable.md "Checks if this reflect.Value can be assigned to a variable.") method.

## Example

```
IMPORT reflect
MAIN
    DEFINE r1, r2 RECORD pkey INTEGER END RECORD
    DEFINE val reflect.Value
    DEFINE s STRING
    LET val = reflect.Value.valueOf(r1)
    DISPLAY "Assigning r1 to r2..."
    CALL val.assignToVariable(r2)
    DISPLAY "Assigning r1 to s..."
    CALL val.assignToVariable(s)
END MAIN
```

Shows:

```
Assigning r1 to r2...
Assigning r1 to s...
Program stopped at 'rec.4gl', line number 10.
FORMS statement error number -8112.
Illegal argument.
```

The second call to the `assignToVariable()` method fails with error -8112 because
the `r1` record cannot be assigned to the simple `STRING` variable
`s`.

## Related links

**Related concepts**  

[reflect.Value.set](4377-reflect-value-set.md "Assigns the specified value to this value object.")

[Data type conversion reference](../08_language-basics/0578-data-type-conversion-reference.md "This topic lists type conversion rules for all data types.")
