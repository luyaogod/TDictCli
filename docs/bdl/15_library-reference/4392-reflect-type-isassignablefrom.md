---
title: "reflect.Type.isAssignableFrom"
source: "fgl-topics/c_fgl_ext_reflect_Type_isAssignableFrom.html"
breadcrumb: "Library reference > Extension packages > The reflect package > The reflect.Type class > reflect.Type methods > reflect.Type.isAssignableFrom"
type: "concept"
---

# reflect.Type.isAssignableFrom

> Checks if this type can be assigned from another type.

## Syntax

```
isAssignableFrom(
     x reflect.Type )
  RETURNS BOOLEAN
```

1. x is the type object to be checked.

## Usage

The `reflect.Type.isAssignableFrom()` method returns `TRUE`, when
the values of the type definition passed as parameter can be assigned to values of the type
represented by this current `reflect.Type` object. If the provided type does not
match this type, the method returns `FALSE`.

## Example

```
IMPORT reflect
TYPE t1 RECORD
        pkey INTEGER,
        name VARCHAR(30)
     END RECORD
TYPE t2 RECORD
        pkey INTEGER,
        name VARCHAR(30)
     END RECORD
TYPE t3 RECORD
        pkey INTEGER,
        name VARCHAR(50),
        addr VARCHAR(100)
     END RECORD
MAIN
    DEFINE rt1 reflect.Type
    DEFINE rt2 reflect.Type
    DEFINE rt3 reflect.Type
    DEFINE v1 t1
    DEFINE v2 t2
    DEFINE v3 t3
    LET rt1 = reflect.Type.typeOf(v1)
    LET rt2 = reflect.Type.typeOf(v2)
    LET rt3 = reflect.Type.typeOf(v3)
    DISPLAY "assignable A : ", rt1.isAssignableFrom(rt1)
    DISPLAY "assignable B : ", rt1.isAssignableFrom(rt2)
    DISPLAY "assignable C : ", rt1.isAssignableFrom(rt3)
END MAIN
```

Shows:

```
assignable A :      1
assignable B :      1
assignable C :      0
```
