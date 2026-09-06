---
title: "STRING.getCharAt"
source: "fgl-topics/c_fgl_datatypes_STRING_getCharAt.html"
breadcrumb: "Library reference > Built-in packages > BDL data types package > STRING data type as class > STRING data type methods > STRING.getCharAt"
type: "concept"
---

# STRING.getCharAt

> Returns the character at the specified position.

## Syntax

```
getCharAt( index INTEGER )
   RETURNS STRING
```

1. index is the position of the character int the string.

## Usage

This method extracts the character at the specified position from the `STRING`
variable.

If the `STRING` variable is `NULL`, or if the
position is out of the bounds of the string, the result will be
`NULL`.

When using byte length semantics, the position is expressed in bytes, and when using char length
semantics, position is specified in characters. In byte length semantics, the method returns
`NULL` if the position does not match a valid character-byte index in the current
string.

> **Important:**
>
> When using [char length semantics](../09_advanced-features/0881-length-semantics-settings.md)
> (FGL\_LENGTH\_SEMANTICS=CHAR), the `getCharAt()` method has the O(N^2) time complexity
> issue: The execution time of this method grows quadratically with the size of the string.

## Example

```
MAIN
  DEFINE s STRING
  LET s = "Some text"
  DISPLAY s.getCharAt(4)
END MAIN
```

Output:

```
e
```

## Related links

**Related concepts**  

[Length semantics settings](../09_advanced-features/0881-length-semantics-settings.md "Length semantics settings")
