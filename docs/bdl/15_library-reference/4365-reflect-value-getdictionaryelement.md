---
title: "reflect.Value.getDictionaryElement"
source: "fgl-topics/c_fgl_ext_reflect_Value_getDictionaryElement.html"
breadcrumb: "Library reference > Extension packages > The reflect package > The reflect.Value class > reflect.Value methods > reflect.Value.getDictionaryElement"
type: "concept"
---

# reflect.Value.getDictionaryElement

> Returns an element of a dictionary.

## Syntax

```
getDictionaryElement(
     key STRING )
  RETURNS reflect.Value
```

1. key is the key whose associated value is to be returned.

## Usage

The `getDictionaryElement()` method returns a `reflect.Value`
object that is a reference to the element associated to the key passed as parameter, in the
dictionary represented by this `reflect.Value` object.

The `reflect.Value` object used to call this method must have been
created from a [`DICTIONARY`](../08_language-basics/0742-dictionaries.md "A dictionary holds an unordered collection of elements accessed by a key.") variable, or is
a `reflect.Value` object returned from a method like [`getField()`](4367-reflect-value-getfield.md "Returns a field by index for a record."), and references a
dictionary.

If the specified key does not exists, a new dictionary element will be inserted, initialized to
`NULL`, and returned. The resulting value object can then be used to set the element
value with [`reflect.Value.set()`](4377-reflect-value-set.md "Assigns the specified value to this value object.").

The new `reflect.Value` object references the original
variable: Any call to a reflection manipulation method modifies the underlying variable
directly.

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
    LET elem = val.getDictionaryElement("key2")
    DISPLAY "val key2   = ", elem.toString()
    CALL elem.set(reflect.Value.copyOf("xxxxx"))
    DISPLAY "dic['key2']= ", dic["key2"]
END MAIN
```

Shows:

```
val key2   = bbbbb
dic['key2']= xxxxx
```
