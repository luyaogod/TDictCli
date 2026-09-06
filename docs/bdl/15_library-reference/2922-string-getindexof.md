---
title: "STRING.getIndexOf"
source: "fgl-topics/c_fgl_datatypes_STRING_getIndexOf.html"
breadcrumb: "Library reference > Built-in packages > BDL data types package > STRING data type as class > STRING data type methods > STRING.getIndexOf"
type: "concept"
---

# STRING.getIndexOf

> Returns the position of a substring.

## Syntax

```
getIndexOf(
     str STRING,
     startIndex INTEGER )
 RETURNS INTEGER
```

1. str is the substring to be searched.
2. startIndex is the starting position for the search.

## Usage

This method scans a `STRING` variable to find the
substring passed as parameter, and returns the position of the substring.

The method starts the search for the substring at the starting position specified by the second
parameter.

The method returns zero if:

- The `STRING` variable is `NULL`.
- The str substring was not found.
- The str substring is `NULL`.
- The start position is out of bounds.

> **Important:**
>
> When using byte length semantics, the position is expressed in bytes,
> and when using char length semantics, it is specified in characters.

## Example

```
MAIN
  DEFINE s STRING
  LET s = "Some text"
  DISPLAY s.getIndexOf("text",1)
END MAIN
```

Output:

```
          6
```

## Related links

**Related concepts**  

[Length semantics settings](../09_advanced-features/0881-length-semantics-settings.md "Length semantics settings")
