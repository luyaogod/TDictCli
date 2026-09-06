---
title: "Example 4: Global variables"
source: "fgl-topics/c_fgl_variables_example_4.html"
breadcrumb: "Language basics > Variables > Examples > Example 4: Global variables"
type: "concept"
description: "This example shows how to define and use global variables. The module \" myglobs.4gl \": GLOBALS DEFINE userid CHAR(20) DEFINE extime DATETIME YEAR TO SECOND END GLOBALS The module \" mylib.4gl \": ..."
---

# Example 4: Global variables

This example shows how to define and use global variables.

The module "myglobs.4gl":

```
GLOBALS
  DEFINE userid CHAR(20)
  DEFINE extime DATETIME YEAR TO SECOND
END GLOBALS
```

The module "mylib.4gl":

```
GLOBALS "myglobs.4gl"

DEFINE s VARCHAR(100)

FUNCTION myfunc()
  DEFINE i INTEGER
  DISPLAY "User Id = " || userid 
  FOR i=1 TO 10
      LET s = "item #" || i 
  END FOR
END FUNCTION
```

The "main.4gl module:

```
GLOBALS "myglobs.4gl"

MAIN
  LET userid = fgl_getenv("LOGNAME")
  LET extime = CURRENT YEAR TO SECOND
  CALL myfunc()
END MAIN
```
