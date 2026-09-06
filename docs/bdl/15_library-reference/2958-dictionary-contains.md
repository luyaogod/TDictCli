---
title: "DICTIONARY.contains"
source: "fgl-topics/c_fgl_DICTIONARY_method_contains.html"
breadcrumb: "Library reference > Built-in packages > BDL data types package > DICTIONARY as class > DICTIONARY methods > DICTIONARY.contains"
type: "concept"
---

# DICTIONARY.contains

> Checks if an element with the given key exists in the dictionary.

## Syntax

```
contains( key STRING )
   RETURNS BOOLEAN
```

1. key is the dictionary key to check.

## Usage

The `contains()` method scans the dictionary to find if there is an element with
the key passed as parameter.

The method returns `TRUE` if the element exists, or `FALSE` if the
element is not found.

## Example

```
MAIN
    DEFINE dict DICTIONARY OF STRING
    LET dict["abc"] = 111
    LET dict["def"] = 222
    LET dict["xyz"] = 999
    DISPLAY dict.contains("def") -- shows 1 (TRUE)
    DISPLAY dict.contains("zzz") -- shows 0 (FALSE)
END MAIN
```
