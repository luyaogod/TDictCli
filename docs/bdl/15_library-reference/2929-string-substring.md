---
title: "STRING.subString"
source: "fgl-topics/c_fgl_datatypes_STRING_subString.html"
breadcrumb: "Library reference > Built-in packages > BDL data types package > STRING data type as class > STRING data type methods > STRING.subString"
type: "concept"
---

# STRING.subString

> Returns a substring from start and end positions in a given string.

## Syntax

```
subString(
     startIndex INTEGER,
     endIndex INTEGER )
   RETURNS STRING
```

1. startIndex is the starting position of the substring.
2. endIndex is the ending position of the substring.

## Usage

This method returns a substring of the current `STRING` variable based on start
and end positions in the original string.

If the `STRING` variable is `NULL`, or when the
positions are out of bounds, the method returns `NULL`.

> **Important:**
>
> When using byte length semantics, the positions are expressed in bytes, and
> when using char length semantics, positions are expressed in characters.

## Example

```
MAIN
  DEFINE s STRING
  LET s = "1234567890"
  DISPLAY s.subString(5,7)
END MAIN
```

Output:

```
567
```

## Related links

**Related concepts**  

[Length semantics settings](../09_advanced-features/0881-length-semantics-settings.md "Length semantics settings")
