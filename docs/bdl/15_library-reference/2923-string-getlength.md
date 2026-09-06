---
title: "STRING.getLength"
source: "fgl-topics/c_fgl_datatypes_STRING_getLength.html"
breadcrumb: "Library reference > Built-in packages > BDL data types package > STRING data type as class > STRING data type methods > STRING.getLength"
type: "concept"
---

# STRING.getLength

> Returns the length of the current string.

## Syntax

```
getLength( )
       RETURNS INTEGER
```

## Usage

This method counts the number of bytes or characters in a `STRING` variable.

Unlike the [`LENGTH`](2787-length.md "Returns the number of characters in a string passed as parameter.")()
function, the `getLength()` method counts the trailing blanks.

If the `STRING` variable is `NULL`, the method
returns zero.

> **Important:**
>
> When using byte length semantics, the length is expressed in bytes, and
> when using char length semantics, it is expressed in characters.

## Example

```
MAIN
  DEFINE s STRING
  LET s = "Some text"
  DISPLAY s.getLength()
END MAIN
```

Output:

```
          9
```

## Related links

**Related concepts**  

[Length semantics settings](../09_advanced-features/0881-length-semantics-settings.md "Length semantics settings")

[STRING.getMultibyteLength](2924-string-getmultibytelength.md "Counts the number of bytes in the string.")
