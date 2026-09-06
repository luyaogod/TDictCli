---
title: "STRING.append"
source: "fgl-topics/c_fgl_datatypes_STRING_append.html"
breadcrumb: "Library reference > Built-in packages > BDL data types package > STRING data type as class > STRING data type methods > STRING.append"
type: "concept"
---

# STRING.append

> Concatenates a string.

## Syntax

```
append( str STRING )
       RETURNS STRING
```

1. str is the string to be concatenated.

## Usage

This method concatenates a string to the current `STRING` variable and
returns the resulting string.

The original `STRING` variable is not modified.

Appending a `NULL` will have no effect: the original string is returned.

## Example

```
MAIN
  DEFINE s STRING
  LET s = "Some text"
  DISPLAY s.append("... more text")
END MAIN
```

Output:

```
Some text... more text
```
