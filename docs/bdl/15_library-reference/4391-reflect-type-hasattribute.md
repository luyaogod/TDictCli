---
title: "reflect.Type.hasAttribute"
source: "fgl-topics/c_fgl_ext_reflect_Type_hasAttribute.html"
breadcrumb: "Library reference > Extension packages > The reflect package > The reflect.Type class > reflect.Type methods > reflect.Type.hasAttribute"
type: "concept"
---

# reflect.Type.hasAttribute

> Indicates if a definition attribute exists in a type.

## Syntax

```
hasAttribute(
     name STRING )
  RETURNS BOOLEAN
```

1. name is the case-sensitive name of the definition attribute to be
   checked.

## Usage

The `hasAttribute()` method returns `TRUE` if the definition
attribute passed as parameter exists for the type referenced by this `reflect.Type`
object. If the attribute is not used, the method returns `FALSE`.

A variable or type definition attribute is specified with the
`ATTRIBUTES()` clause.
> **Note:**
>
> The name of the attribute is case-sensitive.

## Example

```
IMPORT reflect
MAIN
    DEFINE name VARCHAR(30) ATTRIBUTES(json_name="the name")
    DEFINE typ reflect.Type
    LET typ = reflect.Type.typeOf(name)
    DISPLAY "has json_name : ", typ.hasAttribute("json_name")
    DISPLAY "has xxxxx     : ", typ.hasAttribute("xxxxx")
END MAIN
```

Shows:

```
has json_name :      1
has xxxxx     :      0
```
