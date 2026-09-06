---
title: "STRING.toUpperCase"
source: "fgl-topics/c_fgl_datatypes_STRING_toUpperCase.html"
breadcrumb: "Library reference > Built-in packages > BDL data types package > STRING data type as class > STRING data type methods > STRING.toUpperCase"
type: "concept"
---

# STRING.toUpperCase

> Returns the string converted to upper case.

## Syntax

```
toUpperCase( )
       RETURNS STRING
```

## Usage

This method converts the current `STRING` variable to upper case
and returns the resulting string.

If the original `STRING` variable is `NULL`, the
result is `NULL`.

## Example

```
MAIN
  DEFINE s STRING
  LET s = "Some text"
  DISPLAY s.toUpperCase()
END MAIN
```

Output:

```
SOME TEXT
```
