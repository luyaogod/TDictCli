---
title: "reflect.Value.getDictionaryKeys"
source: "fgl-topics/c_fgl_ext_reflect_Value_getDictionaryKeys.html"
breadcrumb: "Library reference > Extension packages > The reflect package > The reflect.Value class > reflect.Value methods > reflect.Value.getDictionaryKeys"
type: "concept"
---

# reflect.Value.getDictionaryKeys

> Returns the keys of a dictionary.

## Syntax

```
getDictionaryKeys( )
  RETURNS DYNAMIC ARRAY OF STRING
```

## Usage

The `getDictionaryKeys()` method returns a `DYNAMIC ARRAY OF
STRING` containing the list of keys defined in the dictionary represented by this
`reflect.Value` object.

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
    DEFINE keys DYNAMIC ARRAY OF STRING
    LET dic["key1"] = "aaaaa"
    LET dic["key2"] = "bbbbb"
    LET val = reflect.Value.valueOf( dic )
    LET keys = val.getDictionaryKeys()
    DISPLAY "keys[1] = ", keys[1]
    DISPLAY "keys[2] = ", keys[2]
END MAIN
```

Shows:

```
keys[1] = key1
keys[2] = key2
```
