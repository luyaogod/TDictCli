---
title: "base.StringBuffer.trimRight"
source: "fgl-topics/c_fgl_ClassStringBuffer_trimRight.html"
breadcrumb: "Library reference > Built-in packages > The base package > The StringBuffer class > base.StringBuffer methods > base.StringBuffer.trimRight"
type: "concept"
---

# base.StringBuffer.trimRight

> Removes trailing blank space (ASCII 32) characters.

## Syntax

```
trimRight()
```

## Usage

The `trimRight()` method removes the trailing blank space characters in the string
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
   CALL buf.trimRight()
   DISPLAY "["||buf.toString()||"]"
END MAIN
```

Output:

```
[  abc]
```

## Related links

**Related concepts**  

[base.StringBuffer.trimRightWhiteSpace](3067-base-stringbuffer-trimrightwhitespace.md "Removes trailing whitespace characters.")
