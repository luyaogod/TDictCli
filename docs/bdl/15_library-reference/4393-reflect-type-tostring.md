---
title: "reflect.Type.toString"
source: "fgl-topics/c_fgl_ext_reflect_Type_toString.html"
breadcrumb: "Library reference > Extension packages > The reflect package > The reflect.Type class > reflect.Type methods > reflect.Type.toString"
type: "concept"
---

# reflect.Type.toString

> Returns the name of a type.

## Syntax

```
toString()
  RETURNS STRING
```

## Usage

The `toString()` method returns the string representation of the type represented
by this `reflect.Type` object.

When the [kind](4388-reflect-type-getkind.md "Returns the kind of a type.") of this type is
`"PRIMITIVE"`, the method returns the raw type definition, such as
`"CHAR(20)", "DATETIME YEAR TO SECOND"`.

When the kind of this type is `"RECORD"` and the type has been defined by a [`TYPE` definition](../08_language-basics/0751-types.md "Types can be defined by the programmer to centralize the definition of complex/structured variables."), the method returns the
name of the type qualified by its module name. For record types defined from a database table, the
method returns `"RECORD LIKE
dbname:tabname.*"`. Otherwise, the method returns
`"RECORD"` for anonymous record types (not defined as a named
`TYPE`).

When the kind is `"ARRAY"` or `"DICTIONARY"`, the method returns
the raw type definition. For example: `DYNAMIC ARRAY OF STRING`, `ARRAY[10,20]
OF MyRecordType`.

When the kind is `"INTERFACE"`, the methdo returns the name of the interface type,
qualified by its module name.

## Example

```
IMPORT reflect
TYPE Person RECORD
        pkey INTEGER,
        name VARCHAR(30)
     END RECORD
TYPE PersonList DYNAMIC ARRAY OF Person
MAIN
    DEFINE arr PersonList
    DISPLAY "type name = ", reflect.Type.typeOf(arr).toString()
END MAIN
```

Shows:

```
type name = DYNAMIC ARRAY OF mymodule.Person
```
