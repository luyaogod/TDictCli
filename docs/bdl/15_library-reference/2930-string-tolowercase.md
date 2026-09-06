---
title: "STRING.toLowerCase"
source: "fgl-topics/c_fgl_datatypes_STRING_toLowerCase.html"
breadcrumb: "Library reference > Built-in packages > BDL data types package > STRING data type as class > STRING data type methods > STRING.toLowerCase"
type: "concept"
---

# STRING.toLowerCase

> Returns the string converted to lower case.

## Syntax

```
toLowerCase( )
       RETURNS STRING
```

## Usage

This method converts the current `STRING` variable to lower case
and returns the resulting string.

If the original `STRING` variable is `NULL`, the
result is `NULL`.

## Example

```
MAIN
  DEFINE s STRING
  LET s = "SOME TEXT"
  DISPLAY s.toLowerCase()
END MAIN
```

Output:

```
some text
```
