---
title: "base.StringBuffer.replace"
source: "fgl-topics/c_fgl_ClassStringBuffer_replace.html"
breadcrumb: "Library reference > Built-in packages > The base package > The StringBuffer class > base.StringBuffer methods > base.StringBuffer.replace"
type: "concept"
---

# base.StringBuffer.replace

> Replace one string with another.

## Syntax

```
replace(
   oldStr STRING,
   newStr STRING,
   occurrences INTEGER )
```

1. oldStr is the string to be replaced.
2. newStr is the new string replacing the old string.
3. occurrences is the number of replacements to do (zero for all).

## Usage

The `replace()` method replaces a string within the current string
buffer with a different string. Specify the original string, replacement string,
and the number of occurrences to replace. Use 0 to replace all occurrences.

## Example

```
MAIN
   DEFINE buf base.StringBuffer
   LET buf = base.StringBuffer.create()
   CALL buf.append("aaxxbbxxcc")
   CALL buf.replace("xx", "zz", 1)
   DISPLAY buf.toString()
END MAIN
```

Output:

```
aazzbbxxcc
```

## Related links

**Related concepts**  

[base.StringBuffer.replaceAt](3057-base-stringbuffer-replaceat.md "Replace part of a string with another string.")
