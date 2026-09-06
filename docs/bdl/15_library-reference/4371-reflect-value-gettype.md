---
title: "reflect.Value.getType"
source: "fgl-topics/c_fgl_ext_reflect_Value_getType.html"
breadcrumb: "Library reference > Extension packages > The reflect package > The reflect.Value class > reflect.Value methods > reflect.Value.getType"
type: "concept"
---

# reflect.Value.getType

> Returns the reflect.Type object of a reflect.Value object.

## Syntax

```
getType()
  RETURNS reflect.Type
```

## Usage

The `getType()` method returns a `reflect.Type` object representing
the type of this `reflect.Value` object.

The type object can then be described with [`reflect.Type` methods](4380-reflect-type-methods.md "Methods for the reflect.Type class.").

## Example

```
IMPORT reflect
MAIN
    DEFINE val reflect.Value
    DEFINE typ reflect.Type
    DEFINE rec RECORD
                   pkey INTEGER,
                   name VARCHAR(30)
               END RECORD
    LET val = reflect.Value.valueOf(rec)
    LET typ = val.getType()
    DISPLAY "type name = ", typ.toString()
END MAIN
```

Shows:

```
type name = RECORD
```

## Related links

**Related concepts**  

[The reflect.Type class](4379-the-reflect-type-class.md "The reflect.Type class is a generic API to inspect types.")
