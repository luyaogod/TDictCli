---
title: "reflect.Value.removeDictionaryElement"
source: "fgl-topics/c_fgl_ext_reflect_Value_removeDictionaryElement.html"
breadcrumb: "Library reference > Extension packages > The reflect package > The reflect.Value class > reflect.Value methods > reflect.Value.removeDictionaryElement"
type: "concept"
---

# reflect.Value.removeDictionaryElement

> Deletes an element of a dictionary.

## Syntax

```
removeDictionaryElement(
     key STRING )
```

1. key is the key of the element to be removed.

## Usage

The `removeDictionaryElement()` method removes an element associated to the
specified key, from the dictionary represented by this `reflect.Value` object.

The `reflect.Value` object used to call this method must have been
created from a [`DICTIONARY`](../08_language-basics/0742-dictionaries.md "A dictionary holds an unordered collection of elements accessed by a key.") variable, or is
a `reflect.Value` object returned from a method like [`getField()`](4367-reflect-value-getfield.md "Returns a field by index for a record."), and references a
dictionary.

If the key does not exist, the method returns silently without error.

## Example

```
IMPORT reflect
MAIN
    DEFINE dic DICTIONARY OF STRING
    DEFINE val reflect.Value
    DEFINE elem reflect.Value
    LET dic["key1"] = "aaaaa"
    LET dic["key2"] = "bbbbb"
    LET val = reflect.Value.valueOf( dic )
    DISPLAY "has key1 : ", val.hasKey("key1")
    DISPLAY "has key2 : ", val.hasKey("key2")
    DISPLAY "has zzzz : ", val.hasKey("zzzz")
END MAIN
```

Shows:

```
has key1 :      1
has key2 :      1
has zzzz :      0
```

## Related links

**Related concepts**  

[DICTIONARY.remove](2962-dictionary-remove.md "Removes an element of the dictionary identified by the key.")
