---
title: "reflect.Value.copyOf"
source: "fgl-topics/c_fgl_ext_reflect_Value_copyOf.html"
breadcrumb: "Library reference > Extension packages > The reflect package > The reflect.Value class > reflect.Value methods > reflect.Value.copyOf"
type: "concept"
---

# reflect.Value.copyOf

> Creates a new reflect.Value object from a copy of an expression.

## Syntax

```
reflect.Value.copyOf(
     val any-type )
  RETURNS reflect.Value
```

1. val is an [expression](../08_language-basics/0592-expressions.md "Shows the possible expressions supported in the language."). This is
   typically a [literal value](../08_language-basics/0585-literals.md "Describes the syntax of literals (constant values) to be used in sources.").
2. any-type can be of a [various kind of types](../09_advanced-features/0847-various-type-specification.md "Some Genero APIs use variant types for parameters or returns.").

## Usage

The `reflect.Value.copyOf()` class method returns a new
`reflect.Value` object, which is a copy of the expression passed as parameter. The
`reflect.Value` object can then be used to call object methods to describe and modify
the original value. Note that certain data types cannot be copied; these exceptions are detailed in
the list below.

When the argument (expression or variable) is defined as a [primitive data type](../08_language-basics/0553-primitive-data-types.md "Selecting the correct data type assists you in the input, storage, and display of your data."), [record structure](../08_language-basics/0715-records.md "Records allow structured program variables definitions.") or [static array](../08_language-basics/0732-static-arrays.md "Static arrays have a predefined and limited size."), the value is cloned to create the
`reflect.Value` object: Any call to a reflection manipulation method will leave the
source unchanged.

Exceptions:

- For [dynamic arrays](../08_language-basics/0734-dynamic-arrays.md), [dictionaries](../08_language-basics/0742-dictionaries.md "A dictionary holds an unordered collection of elements accessed by a key.") and objects of [built-in
  classes](2908-built-in-packages.md "These topics cover the built-in classes provided by the Genero Business Development Language.") the `reflect.Value.copyOf()` makes a copy of the reference to the
  variable: Any call to a reflection manipulation method will change the source.

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
    LET val = reflect.Value.copyOf( rec )
    LET rec.pkey = 102
    LET rec.name = "Jessica PARS"
    DISPLAY "val pkey = ", val.getField(1).toString()
    DISPLAY "val name = ", val.getField(2).toString()
END MAIN
```

Shows:

```
val pkey = 101
val name = Mike FITZPATRICK
```

## Related links

**Related concepts**  

[reflect.Value.valueOf](4357-reflect-value-valueof.md "Creates a new reflect.Value object as a reference to the original variable.")
