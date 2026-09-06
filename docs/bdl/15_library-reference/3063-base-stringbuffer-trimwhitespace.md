---
title: "base.StringBuffer.trimWhiteSpace"
source: "fgl-topics/c_fgl_ClassStringBuffer_trimWhiteSpace.html"
breadcrumb: "Library reference > Built-in packages > The base package > The StringBuffer class > base.StringBuffer methods > base.StringBuffer.trimWhiteSpace"
type: "concept"
---

# base.StringBuffer.trimWhiteSpace

> Remove leading and trailing whitespace characters.

## Syntax

```
trimWhiteSpace()
```

## Usage

The `trimWhiteSpace()` method removes the leading and trailing whitespace
characters in the string buffer.

The method considers as whitespace characters all characters less than or equal to blank space
(`ASCII(32)`). This includes tab (`\t`), newline
(`\n`), carriage-return (`\r`) and form-feed
(`\f`).

## Example

```
MAIN
   DEFINE buf base.StringBuffer
   LET buf = base.StringBuffer.create()
   CALL buf.append("\n\t  abc  \n\t")
   CALL buf.trimWhiteSpace()
   DISPLAY "["||buf.toString()||"]" -- Shows [abc]
END MAIN
```

Output:

```
[abc]
```

## Related links

**Related concepts**  

[base.StringBuffer.trim](3062-base-stringbuffer-trim.md "Remove leading and trailing blank space (ASCII 32) characters.")
