---
title: "base.StringBuffer.trimLeftWhiteSpace"
source: "fgl-topics/c_fgl_ClassStringBuffer_trimLeftWhiteSpace.html"
breadcrumb: "Library reference > Built-in packages > The base package > The StringBuffer class > base.StringBuffer methods > base.StringBuffer.trimLeftWhiteSpace"
type: "concept"
---

# base.StringBuffer.trimLeftWhiteSpace

> Removes leading whitespace characters.

## Syntax

```
trimLeftWhiteSpace()
```

## Usage

The `trimLeftWhiteSpace()` method removes the leading whitespace characters in the
string buffer.

The method considers as whitespace characters all characters less than or equal to blank space
(`ASCII(32)`). This includes tab (`\t`), newline
(`\n`), carriage-return (`\r`) and form-feed
(`\f`).

## Example

```
MAIN
   DEFINE buf base.StringBuffer
   LET buf = base.StringBuffer.create()
   CALL buf.append("\n\t  abc  ")
   CALL buf.trimLeftWhiteSpace()
   DISPLAY "["||buf.toString()||"]"
END MAIN
```

Output:

```
[abc  ]
```

## Related links

**Related concepts**  

[base.StringBuffer.trimLeft](3064-base-stringbuffer-trimleft.md "Removes leading blank space (ASCII 32) characters.")
