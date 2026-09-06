---
title: "STRING.equalsIgnoreCase"
source: "fgl-topics/c_fgl_datatypes_STRING_equalsIgnoreCase.html"
breadcrumb: "Library reference > Built-in packages > BDL data types package > STRING data type as class > STRING data type methods > STRING.equalsIgnoreCase"
type: "concept"
---

# STRING.equalsIgnoreCase

> Makes a case-insensitive string comparison.

## Syntax

```
equalsIgnoreCase( str STRING )
       RETURNS BOOLEAN
```

1. str is the string to compare with.

## Usage

This method compares a string to the current `STRING` variable by ignoring the
character case, and returns `TRUE` if both strings match or when both strings are
`NULL`.

## Example

```
MAIN
  DEFINE s STRING
  LET s = "white"
  IF s.equalsIgnoreCase("WHITE") THEN
     DISPLAY "Matches"
  END IF
END MAIN
```
