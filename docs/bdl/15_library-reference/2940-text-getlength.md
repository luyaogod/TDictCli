---
title: "TEXT.getLength"
source: "fgl-topics/c_fgl_datatypes_TEXT_getLength.html"
breadcrumb: "Library reference > Built-in packages > BDL data types package > TEXT data type as class > TEXT data type methods > TEXT.getLength"
type: "concept"
---

# TEXT.getLength

> Returns the length of TEXT content.

## Syntax

```
getLength( )
   RETURNS INTEGER
```

## Usage

This method returns the number of bytes in `TEXT` data.

This method always returns a number of bytes, even when using character length semantics.

## Example

```
MAIN
  DEFINE t TEXT
  LOCATE t IN MEMORY
  LET t = "aaaaaaaaaaaaa"
  DISPLAY t.getLength() -- Shows 13
END MAIN
```

## Related links

**Related concepts**  

[Length semantics settings](../09_advanced-features/0881-length-semantics-settings.md "Length semantics settings")
