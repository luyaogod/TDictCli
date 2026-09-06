---
title: "Example 2: Escaped delimiters"
source: "fgl-topics/c_fgl_ClassStringTokenizer_example_2.html"
breadcrumb: "Library reference > Built-in packages > The base package > The StringTokenizer class > Examples > Example 2: Escaped delimiters"
type: "concept"
description: "The next example uses the pipe as delimiter, ans specifies the backslash as escape character. Note that backslash has to be doubled in string literals . Program code: MAIN DEFINE tok ..."
---

# Example 2: Escaped delimiters

The next example uses the pipe as delimiter, ans specifies the backslash as escape character.
Note that backslash has to be doubled in [string
literals](../08_language-basics/0588-text-literals.md "Text literals define a character string in an expression.").

Program code:

```
MAIN
  DEFINE tok base.StringTokenizer
  DEFINE x INTEGER
  DEFINE s STRING
  LET tok = base.StringTokenizer.createExt("||\\|aaa||bbc|","|","\\",FALSE)
  WHILE tok.hasMoreTokens()
    LET s = tok.nextToken()
    DISPLAY (x:=x+1)," [", s, "]"
  END WHILE
END MAIN
```

Output:

```
          1 [|aaa]
          2 [bbc]
```
