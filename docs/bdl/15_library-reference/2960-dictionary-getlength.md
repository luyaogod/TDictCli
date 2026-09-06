---
title: "DICTIONARY.getLength"
source: "fgl-topics/c_fgl_DICTIONARY_method_getLength.html"
breadcrumb: "Library reference > Built-in packages > BDL data types package > DICTIONARY as class > DICTIONARY methods > DICTIONARY.getLength"
type: "concept"
---

# DICTIONARY.getLength

> Returns the number of elements in the dictionary.

## Syntax

```
getLength( )
   RETURNS INTEGER
```

## Usage

The `getLength()` method returns the total number of elements in the
dictionary.

## Example

```
MAIN
    DEFINE dict DICTIONARY OF STRING
    LET dict["abc"] = 111
    LET dict["def"] = 222
    LET dict["xyz"] = 999
    DISPLAY dict.getLength()   -- shows 3
END MAIN
```
