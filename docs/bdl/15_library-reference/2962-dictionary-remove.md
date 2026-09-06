---
title: "DICTIONARY.remove"
source: "fgl-topics/c_fgl_DICTIONARY_method_remove.html"
breadcrumb: "Library reference > Built-in packages > BDL data types package > DICTIONARY as class > DICTIONARY methods > DICTIONARY.remove"
type: "concept"
---

# DICTIONARY.remove

> Removes an element of the dictionary identified by the key.

## Syntax

```
remove( key STRING )
```

1. key is the dictionary key of the element to remove.

## Usage

The `remove()` method deletes and element of the dictionary, identified by the key
passed as paramater.

## Example

```
MAIN
    DEFINE dict DICTIONARY OF STRING
    LET dict["abc"] = 111
    CALL dict.remove("abc")
    DISPLAY dict.contains("abc") -- shows 0 (FALSE)
END MAIN
```
