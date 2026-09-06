---
title: "Example 2: Modify a StringBuffer with a function"
source: "fgl-topics/c_fgl_ClassStringBuffer_example_2.html"
breadcrumb: "Library reference > Built-in packages > The base package > The StringBuffer class > Examples > Example 2: Modify a StringBuffer with a function"
type: "concept"
description: "MAIN DEFINE buf base.StringBuffer LET buf = base.StringBuffer.create() CALL modify(buf) DISPLAY \"buf is \", buf.toString() END MAIN FUNCTION modify(sb) DEFINE sb base.StringBuffer CALL ..."
---

# Example 2: Modify a StringBuffer with a function

```
MAIN
  DEFINE buf base.StringBuffer 
  LET buf = base.StringBuffer.create()
  CALL modify(buf)
  DISPLAY "buf is ", buf.toString()
END MAIN

FUNCTION modify(sb)
  DEFINE sb base.StringBuffer 
  CALL sb.append("more")
  DISPLAY "sb is ", sb.toString()
END FUNCTION
```

Output:

```
sb is more
buf is more
```
