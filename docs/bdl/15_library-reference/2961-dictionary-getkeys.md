---
title: "DICTIONARY.getKeys"
source: "fgl-topics/c_fgl_DICTIONARY_method_getKeys.html"
breadcrumb: "Library reference > Built-in packages > BDL data types package > DICTIONARY as class > DICTIONARY methods > DICTIONARY.getKeys"
type: "concept"
---

# DICTIONARY.getKeys

> Returns a dynamic array of all keys of the dictionary.

## Syntax

```
getKeys( )
   RETURNS DYNAMIC ARRAY OF STRING
```

## Usage

The `getKeys()` method builds and returns a `DYNAMIC ARRAY OF
STRING` with all keys contained in the dictionary.

It is then possible to scan the dictionary with the keys returned by this method.

A dictionary is an unordered list of elements. The keys might not be returned in the same order
as they have been added to the dictionary.

## Example

```
MAIN
    DEFINE dict DICTIONARY OF STRING,
           keys DYNAMIC ARRAY OF STRING,
           i INTEGER
    LET dict["first"]  = "abc"
    LET dict["second"] = "def"
    LET dict["lasy"]   = "xyz"
    LET keys = dict.getKeys()
    FOR i = 1 to keys.getLength()
        DISPLAY i, dict[keys[i]]
    END FOR
END MAIN
```
