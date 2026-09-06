---
title: "reflect.Type.getAttribute"
source: "fgl-topics/c_fgl_ext_reflect_Type_getAttribute.html"
breadcrumb: "Library reference > Extension packages > The reflect package > The reflect.Type class > reflect.Type methods > reflect.Type.getAttribute"
type: "concept"
---

# reflect.Type.getAttribute

> Returns the value of a definition attribute.

## Syntax

```
getAttribute(
     name STRING )
  RETURNS STRING
```

1. name is the case-sensitive name of the definition attribute to be
   returned.

## Usage

The `getAttribute()` method returns the value associated to a [definition attribute](../08_language-basics/0692-attributes-on-variable-definitions.md "Variables can be defined with meta-data information.") for this
`reflect.Type` object representing a type.

If the specified attribute is not defined for this type object, the
`getAttribute()` method returns `NULL`.

A variable or type definition attribute is specified with the
`ATTRIBUTES()` clause.
> **Note:**
>
> The name of the attribute is case-sensitive.

## Example

```
IMPORT reflect
MAIN
    DEFINE rec RECORD ATTRIBUTES(json_name="a person")
               pkey INTEGER ATTRIBUTES(json_name="the key"),
               name VARCHAR(30) ATTRIBUTES(json_name="the name"),
               addr VARCHAR(100)
           END RECORD
    DEFINE typ reflect.Type
    LET typ = reflect.Type.typeOf( rec )
    DISPLAY "rec json_name     : ", typ.getAttribute("json_name")
    DISPLAY "field 1 json_name : ", typ.getFieldType(1).getAttribute("json_name")
    DISPLAY "field 2 json_name : ", typ.getFieldType(2).getAttribute("json_name")
    DISPLAY "field 3 json_name : ", typ.getFieldType(3).getAttribute("json_name")
END MAIN
```

Shows:

```
rec json_name     : a person
field 1 json_name : the key
field 2 json_name : the name
field 3 json_name :
```
