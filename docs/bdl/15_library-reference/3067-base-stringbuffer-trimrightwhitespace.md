---
title: "base.StringBuffer.trimRightWhiteSpace"
source: "fgl-topics/c_fgl_ClassStringBuffer_trimRightWhiteSpace.html"
breadcrumb: "Library reference > Built-in packages > The base package > The StringBuffer class > base.StringBuffer methods > base.StringBuffer.trimRightWhiteSpace"
type: "concept"
---

# base.StringBuffer.trimRightWhiteSpace

> Removes trailing whitespace characters.

## Syntax

```
trimRightWhiteSpace()
```

## Usage

The `trimRightWhiteSpace()` method removes the trailing whitespace characters in
the string buffer.

The method considers as whitespace characters all characters less than or equal to blank space
(`ASCII(32)`). This includes tab (`\t`), newline
(`\n`), carriage-return (`\r`) and form-feed
(`\f`).

## Example

```
MAIN
   DEFINE buf base.StringBuffer
   LET buf = base.StringBuffer.create()
   CALL buf.append("  abc  \n\t")
   CALL buf.trimRightWhiteSpace()
   DISPLAY "["||buf.toString()||"]"
END MAIN
```

Output:

```
[  abc]
```

## Related links

**Related concepts**  

[base.StringBuffer.trimRight](3066-base-stringbuffer-trimright.md "Removes trailing blank space (ASCII 32) characters.")
