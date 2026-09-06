---
title: "Example 3: Empty tokens"
source: "fgl-topics/c_fgl_ClassStringTokenizer_example_3.html"
breadcrumb: "Library reference > Built-in packages > The base package > The StringTokenizer class > Examples > Example 3: Empty tokens"
type: "concept"
description: "This example shows how to take empty tokens into account, Program code: MAIN DEFINE tok base.StringTokenizer DEFINE x INTEGER DEFINE s STRING LET tok = ..."
---

# Example 3: Empty tokens

This example shows how to take empty tokens into account,

Program code:

```
MAIN
  DEFINE tok base.StringTokenizer
  DEFINE x INTEGER
  DEFINE s STRING
  LET tok = base.StringTokenizer.createExt("|a|bb||ccc|","|",NULL,TRUE)
  WHILE tok.hasMoreTokens()
    LET s = tok.nextToken()
    DISPLAY (x:=x+1)," [", s, "]",
      COLUMN 30, s.getLength(), " IS NULL: ", (s IS NULL)
  END WHILE
END MAIN
```

Output:

```
          1 []                         0 IS NULL:      0
          2 [a]                        1 IS NULL:      0
          3 [bb]                       2 IS NULL:      0
          4 []                         0 IS NULL:      0
          5 [ccc]                      3 IS NULL:      0
          6 []                         0 IS NULL:      0
```
