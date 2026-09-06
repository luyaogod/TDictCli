---
title: "reflect.Value.isNull"
source: "fgl-topics/c_fgl_ext_reflect_Value_isNull.html"
breadcrumb: "Library reference > Extension packages > The reflect package > The reflect.Value class > reflect.Value methods > reflect.Value.isNull"
type: "concept"
---

# reflect.Value.isNull

> Checks if this reflect.Value is NULL.

## Syntax

```
isNull( )
  RETURNS BOOLEAN
```

## Usage

The `isNull()` method returns `TRUE`, if the
`reflect.Value` object references a value that is `NULL`.

## Example

```
IMPORT reflect
MAIN
    DEFINE s STRING
    LET s = "abc"
    IF reflect.Value.valueOf(s).isNull() THEN
        DISPLAY "s is null"
    ELSE
        DISPLAY "s is not null"
    END IF
END MAIN
```

Shows:

```
s is not null
```

## Related links

**Related concepts**  

[Data type conversion reference](../08_language-basics/0578-data-type-conversion-reference.md "This topic lists type conversion rules for all data types.")

[Formatting data](../08_language-basics/0580-formatting-data.md "Explains data to string conversion options of the language.")
