---
title: "STRING.expandTabs"
source: "fgl-topics/c_fgl_datatypes_STRING_expandTabs.html"
breadcrumb: "Library reference > Built-in packages > BDL data types package > STRING data type as class > STRING data type methods > STRING.expandTabs"
type: "concept"
---

# STRING.expandTabs

> Converts TAB characters to a 8 space characters.

## Syntax

```
expandTabs( )
       RETURNS STRING
```

## Usage

This method finds all TAB (ASCII 9) characters in the string and replaces each TAB character by 8
space (ASCII 32) characters. The new string is then returned by the method.

## Example

```
MAIN
    DEFINE s STRING
    LET s = "a\tb\tc"
    DISPLAY s
    DISPLAY s.expandTabs()
END MAIN
```

Output:

```
a	b	c
a       b       c
```

First line contains 2 TAB characters, while second line contains
2x8 space characters.
