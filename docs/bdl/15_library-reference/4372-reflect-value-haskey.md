---
title: "reflect.Value.hasKey"
source: "fgl-topics/c_fgl_ext_reflect_Value_hasKey.html"
breadcrumb: "Library reference > Extension packages > The reflect package > The reflect.Value class > reflect.Value methods > reflect.Value.hasKey"
type: "concept"
---

# reflect.Value.hasKey

> Checks if existence of an element in a dictionary.

## Syntax

```
hasKey(
     key STRING )
  RETURNS BOOLEAN
```

1. key is the key to be checked.

## Usage

The `hasKey()` method returns `TRUE`, if the
`reflect.Value` object is a `DICTIONARY` that contains an element
identified with the specified key.

The `reflect.Value` object used to call this method must have been
created from a [`DICTIONARY`](../08_language-basics/0742-dictionaries.md "A dictionary holds an unordered collection of elements accessed by a key.") variable, or is
a `reflect.Value` object returned from a method like [`getField()`](4367-reflect-value-getfield.md "Returns a field by index for a record."), and references a
dictionary.

## Example

```
IMPORT reflect
MAIN
    DEFINE dic DICTIONARY OF STRING
    DEFINE val reflect.Value
    LET dic["key1"] = "aaaaa"
    LET dic["key2"] = "bbbbb"
    LET val = reflect.Value.valueOf( dic )
    CALL val.removeDictionaryElement("key1")
    DISPLAY "has key1 : ", val.hasKey("key1")
END MAIN
```

Shows:

```
has key1 :      0
```

## Related links

**Related concepts**  

[DICTIONARY.contains](2958-dictionary-contains.md "Checks if an element with the given key exists in the dictionary.")
