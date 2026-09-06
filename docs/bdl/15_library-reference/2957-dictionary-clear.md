---
title: "DICTIONARY.clear"
source: "fgl-topics/c_fgl_DICTIONARY_method_clear.html"
breadcrumb: "Library reference > Built-in packages > BDL data types package > DICTIONARY as class > DICTIONARY methods > DICTIONARY.clear"
type: "concept"
---

# DICTIONARY.clear

> Removes all elements of the dictionary.

## Syntax

```
clear( )
```

## Usage

The `clear()` method cleans up the dictionary, by removing all its elements.

For example, if the dictionary is not empty, use the `clear()` method just before
filling the dictionary with a new set of elements.

## Example

```
FUNCTION fill_dictionary(dict)
    DEFINE dict DICTIONARY OF STRING
    CALL dict.clear()
    LET dict["abc"] = 111
    LET dict["def"] = 222
    LET dict["xyz"] = 999
END FUNCTION
```
