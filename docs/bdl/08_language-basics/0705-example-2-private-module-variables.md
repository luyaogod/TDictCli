---
title: "Example 2: PRIVATE module variables"
source: "fgl-topics/c_fgl_variables_example_2.html"
breadcrumb: "Language basics > Variables > Examples > Example 2: PRIVATE module variables"
type: "concept"
description: "This example shows how to define a private module variable. PRIVATE DEFINE s VARCHAR(100) FUNCTION myfunc() DEFINE i INTEGER FOR i=1 TO 10 LET s = \"item #\" || i END FOR END FUNCTION"
---

# Example 2: PRIVATE module variables

This example shows how to define a private module variable.

```
PRIVATE DEFINE s VARCHAR(100)

FUNCTION myfunc()
  DEFINE i INTEGER
  FOR i=1 TO 10
      LET s = "item #" || i 
  END FOR
END FUNCTION
```
