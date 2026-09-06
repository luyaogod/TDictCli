---
title: "base.StringBuffer.trim"
source: "fgl-topics/c_fgl_ClassStringBuffer_trim.html"
breadcrumb: "Library reference > Built-in packages > The base package > The StringBuffer class > base.StringBuffer methods > base.StringBuffer.trim"
type: "concept"
---

# base.StringBuffer.trim

> Remove leading and trailing blank space (ASCII 32) characters.

## Syntax

```
trim()
```

## Usage

The `trim()` method removes the leading and trailing blank space characters in the
string buffer.

The method removes only blank space (`ASCII(32)`) characters. Characters tab
(`\t`), newline (`\n`), carriage-return (`\r`) and
form-feed (`\f`) are not removed.

## Example

```
MAIN
   DEFINE buf base.StringBuffer
   LET buf = base.StringBuffer.create()
   CALL buf.append("  abc  ")
   CALL buf.trim()
   DISPLAY "["||buf.toString()||"]" -- Shows [abc]
END MAIN
```

Output:

```
[abc]
```

## Related links

**Related concepts**  

[base.StringBuffer.trimWhiteSpace](3063-base-stringbuffer-trimwhitespace.md "Remove leading and trailing whitespace characters.")
