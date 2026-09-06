---
title: "STRING.equals"
source: "fgl-topics/c_fgl_datatypes_STRING_equals.html"
breadcrumb: "Library reference > Built-in packages > BDL data types package > STRING data type as class > STRING data type methods > STRING.equals"
type: "concept"
---

# STRING.equals

> Compares a string to the content of a string variable.

## Syntax

```
equals( str STRING )
       RETURNS BOOLEAN
```

1. str is the string to compare with.

## Usage

This method compares a string to the current `STRING` variable and returns
`TRUE` if both strings match or when both strings are `NULL`

## Example

```
MAIN
  DEFINE s STRING
  LET s = "ab"
  DISPLAY '"ab"  s.equals("ab") :', s.equals("ab")
  DISPLAY '"ab"  s.equals(NULL) :', s.equals(NULL)
  LET s = NULL;
  DISPLAY 'NULL  s.equals("ab") :', s.equals("ab")
  DISPLAY 'NULL  s.equals(NULL) :', s.equals(NULL)
END MAIN
```

Output:

```
"ab"  s.equals("ab") :          1
"ab"  s.equals(NULL) :          0
NULL  s.equals("ab") :          0
NULL  s.equals(NULL) :          1
```
