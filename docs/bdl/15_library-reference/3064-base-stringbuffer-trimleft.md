---
title: "base.StringBuffer.trimLeft"
source: "fgl-topics/c_fgl_ClassStringBuffer_trimLeft.html"
breadcrumb: "Library reference > Built-in packages > The base package > The StringBuffer class > base.StringBuffer methods > base.StringBuffer.trimLeft"
type: "concept"
---

# base.StringBuffer.trimLeft

> Removes leading blank space (ASCII 32) characters.

## Syntax

```
trimLeft()
```

## Usage

The `trimLeft()` method removes the leading blank space characters in the string
buffer.

The method removes only blank space (`ASCII(32)`) characters. Characters tab
(`\t`), newline (`\n`), carriage-return (`\r`) and
form-feed (`\f`) are not removed.

## Example

```
MAIN
   DEFINE buf base.StringBuffer
   LET buf = base.StringBuffer.create()
   CALL buf.append("  abc  ")
   CALL buf.trimLeft()
   DISPLAY "["||buf.toString()||"]"
END MAIN
```

Output:

```
[abc  ]
```

## Related links

**Related concepts**  

[base.StringBuffer.trimLeftWhiteSpace](3065-base-stringbuffer-trimleftwhitespace.md "Removes leading whitespace characters.")
