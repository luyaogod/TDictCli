---
title: "reflect.Value.valueOf"
source: "fgl-topics/c_fgl_ext_reflect_Value_valueOf.html"
breadcrumb: "Library reference > Extension packages > The reflect package > The reflect.Value class > reflect.Value methods > reflect.Value.valueOf"
type: "concept"
---

# reflect.Value.valueOf

> Creates a new reflect.Value object as a reference to the original variable.

## Syntax

```
reflect.Value.valueOf(
     val any-type )
  RETURNS reflect.Value
```

1. val must be a [variable](../08_language-basics/0686-variables.md "Explains how to define program variables."), it
   cannot be an expression or literal.
   > **Note:**
   >
   > The val parameter is passed by reference
   > to the method.
2. any-type can be of a [various kind of types](../09_advanced-features/0847-various-type-specification.md "Some Genero APIs use variant types for parameters or returns."), except [`STRING`](../08_language-basics/0567-string.md "The STRING data type is a variable-length, dynamically allocated character string data type, without limitation.").
   > **Note:**
   >
   > The type of the variable cannot be `STRING`,
   > use `copyOf()` instead.

## Usage

The `reflect.Value.valueOf()` method returns a new `reflect.Value`
object, which is a reference to the variable passed as parameter. The `reflect.Value`
object can then be used to call object methods to describe and modify the value. Note that certain
data types cannot be referenced; these exceptions are detailed in the list below.

The new `reflect.Value` object references the original
variable: Any call to a reflection manipulation method modifies the underlying variable
directly.

Exceptions:

- The variable passed as parameter to the `valueOf()` method cannot be of type
  `STRING`. To create a `reflect.Value` object from a
  `STRING` variable, use the [`copyOf()`](4356-reflect-value-copyof.md "Creates a new reflect.Value object from a copy of an expression.") method.
- When passing a variable defined with a [primitive
  type](../08_language-basics/0553-primitive-data-types.md "Selecting the correct data type assists you in the input, storage, and display of your data."), the `reflect.Value.valueOf()` method returns a copy (like [`reflect.Value.copyOf()`](4356-reflect-value-copyof.md "Creates a new reflect.Value object from a copy of an expression.")), instead
  of a reference. As result, changing the returned `reflect.Value` object will have no
  effect on the original variable.

## Example

```
IMPORT reflect
MAIN
    DEFINE rec RECORD
               pkey INTEGER,
               name VARCHAR(50)
           END RECORD
    DEFINE val reflect.Value
    LET rec.pkey = 101
    LET rec.name = "Mike FITZPATRICK"
    LET val = reflect.Value.valueOf( rec )
    CALL val.getField(1).set(reflect.Value.copyOf(999))
    CALL val.getField(2).set(reflect.Value.copyOf("Scott PARS"))
    DISPLAY "rec pkey = ", rec.pkey
    DISPLAY "rec name = ", rec.name
END MAIN
```

Shows:

```
rec pkey =         999
rec name = Scott PARS
```

## Related links

**Related concepts**  

[reflect.Value.copyOf](4356-reflect-value-copyof.md "Creates a new reflect.Value object from a copy of an expression.")
