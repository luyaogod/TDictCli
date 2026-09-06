---
title: "Example 1: Add strings to a StringBuffer"
source: "fgl-topics/c_fgl_ClassStringBuffer_example_1.html"
breadcrumb: "Library reference > Built-in packages > The base package > The StringBuffer class > Examples > Example 1: Add strings to a StringBuffer"
type: "concept"
description: "MAIN DEFINE buf base.StringBuffer LET buf = base.StringBuffer.create() CALL buf.append(\"abc\") DISPLAY buf.toString() CALL buf.append(\"def\") DISPLAY buf.toString() CALL buf.append(123456) DISPLAY ..."
---

# Example 1: Add strings to a StringBuffer

```
MAIN
  DEFINE buf base.StringBuffer 
  LET buf = base.StringBuffer.create()
  CALL buf.append("abc")
  DISPLAY buf.toString()
  CALL buf.append("def")
  DISPLAY buf.toString()
  CALL buf.append(123456)
  DISPLAY buf.toString()
END MAIN
```

Output:

```
abc
abcdef
abcdef123456
```
