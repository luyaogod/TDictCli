---
title: "base.StringBuffer.replaceAt"
source: "fgl-topics/c_fgl_ClassStringBuffer_replaceAt.html"
breadcrumb: "Library reference > Built-in packages > The base package > The StringBuffer class > base.StringBuffer methods > base.StringBuffer.replaceAt"
type: "concept"
---

# base.StringBuffer.replaceAt

> Replace part of a string with another string.

## Syntax

```
replaceAt(
   index INTEGER,
   length INTEGER,
   str STRING )
```

1. index is position where the replacement starts.
2. length is the number of characters to be replaced.
3. str is the replacement string.

## Usage

The `replaceAt()` method replaces part of the current string
with another string.

The parameters are integers indicating the position at which the replacement
should start, the number of characters to be replaced, and the replacement string.

The first position in the string is 1.

> **Important:**
>
> When using byte length semantics, the position and length are expressed in
> bytes. When using [char length semantics](../09_advanced-features/0881-length-semantics-settings.md), the unit is
> characters. This matters when using a multibyte locale such as UTF-8.

## Example

```
MAIN
   DEFINE buf base.StringBuffer
   LET buf = base.StringBuffer.create()
   CALL buf.append("abxxxxef")
   CALL buf.replaceAt(3,4,"cd")
   DISPLAY buf.toString()
END MAIN
```

Output:

```
abcdef
```

## Related links

**Related concepts**  

[base.StringBuffer.replace](3056-base-stringbuffer-replace.md "Replace one string with another.")
