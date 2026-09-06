---
title: "reflect.Value.clear"
source: "fgl-topics/c_fgl_ext_reflect_Value_clear.html"
breadcrumb: "Library reference > Extension packages > The reflect package > The reflect.Value class > reflect.Value methods > reflect.Value.clear"
type: "concept"
---

# reflect.Value.clear

> Removes all elements of the reflect.Value referencing a collection.

## Syntax

```
clear( )
```

## Usage

The `clear()` method clears the dynamic array or
dictionary referenced by this `reflect.Value` object.

The
`reflect.Value` object used to call this method must have been created from a [`DYNAMIC ARRAY`](../08_language-basics/0734-dynamic-arrays.md), a [`DICTIONARY`](../08_language-basics/0742-dictionaries.md "A dictionary holds an unordered collection of elements accessed by a key.") variable, or is a
`reflect.Value` object returned from a method like [`getField()`](4367-reflect-value-getfield.md "Returns a field by index for a record."), and references a
dynamic array or dictionary.

## Example

```
IMPORT reflect
MAIN
    DEFINE val reflect.Value
    DEFINE arr DYNAMIC ARRAY OF STRING
    LET arr[1] = "elem1"
    LET arr[2] = "elem2"
    LET val = reflect.Value.valueOf(arr)
    CALL val.clear()
    DISPLAY "arr len =", arr.getLength()
END MAIN
```

Shows:

```
arr len =          0
```

## Related links

**Related concepts**  

[Data type conversion reference](../08_language-basics/0578-data-type-conversion-reference.md "This topic lists type conversion rules for all data types.")

[Formatting data](../08_language-basics/0580-formatting-data.md "Explains data to string conversion options of the language.")
